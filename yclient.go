package dorfyn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// queryParams is a map of query parameters, mapping from parameter name to value.
type queryParams map[string]string

// yClient is a Yahoo Finance client.
type yClient struct {
	expiry  time.Time
	cookies string
	crumb   string
}

const (
	defaultHTTPTimeout = 80 * time.Second
	yFinURL            = "https://query1.finance.yahoo.com"

	crumbURL  = yFinURL + "/v1/test/getcrumb"
	cookieURL = "https://login.yahoo.com"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0"

	enableDump = true
)

var (
	httpClient = &http.Client{Timeout: defaultHTTPTimeout}
)

// fetchCookies fetches cookies from Yahoo Finance.
// The cookies are required to fetch the crumb that is in turn required to fetch quotes.
func fetchCookies() (string, time.Time, error) {
	logInfo("Fetching cookies...")

	request, err := http.NewRequest("GET", cookieURL, nil)
	if err != nil {
		logError("Can't create cookie request: %v\n", err)
		return "", time.Time{}, err
	}

	request.Header = http.Header{
		"Accept":                   {"*/*"},
		"Accept-Encoding":          {"gzip, deflate, br"},
		"Accept-Language":          {"en-US,en;q=0.5"},
		"Connection":               {"keep-alive"},
		"Host":                     {"login.yahoo.com"},
		"Sec-Fetch-Dest":           {"document"},
		"Sec-Fetch-Mode":           {"navigate"},
		"Sec-Fetch-Site":           {"none"},
		"Sec-Fetch-User":           {"?1"},
		"TE":                       {"trailers"},
		"Update-Insecure-Requests": {"1"},
		"User-Agent":               {userAgent},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		logError("Can't fetch cookies: %v\n", err)
		return "", time.Time{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logError("Can't close cookie response body: %v\n", err)
		}
	}(response.Body)

	var result string

	// Default expiry is ten years in the future
	var expiry = time.Now().AddDate(10, 0, 0)

	for _, cookie := range response.Cookies() {

		logDebug("Considering cookie: %v\n", cookie)

		if cookie.MaxAge <= 0 || cookie.Name == "AS" {
			logDebug("Cookie ignored")
			continue
		}

		logDebug("Cookie accepted")

		cookieExpiry := time.Now().Add(time.Duration(cookie.MaxAge) * time.Second)
		result += cookie.Name + "=" + cookie.Value + "; "

		// set expiry to the latest cookie expiry if smaller than the current expiry
		if cookie.Expires.Before(cookieExpiry) {
			logDebug("Setting expiry to %v\n", cookieExpiry)
			expiry = cookieExpiry
		}
	}
	result = strings.TrimSuffix(result, "; ")
	return result, expiry, nil
}

// fetchCrumb fetches a crumb from Yahoo Finance. The crumb is required to fetch quotes.
func fetchCrumb(cookies string) (string, error) {
	logInfo("Fetching crumb with cookies: %s\n", cookies)
	request, err := http.NewRequest("GET", crumbURL, nil)
	if err != nil {
		logError("Can't create crumb request: %v\n", err)
		return "", err
	}

	request.Header = http.Header{
		"Accept":          {"*/*"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Connection":      {"keep-alive"},
		"Content-Type":    {"text/plain"},
		"Cookie":          {cookies},
		"Host":            {"query1.finance.yahoo.com"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-site"},
		"TE":              {"trailers"},
		"User-Agent":      {userAgent},
	}

	response, err := httpClient.Do(request)
	if err != nil {
		logError("Can't fetch crumb: %v\n", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logError("Can't close crumb response body: %v\n", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logError("Can't read crumb response: %v\n", err)
		return "", err
	}

	return string(body[:]), nil
}

// refreshCrumb refreshes the cookie and crumb.
func (client *yClient) refreshCrumb() error {
	logInfo("Refreshing crumb...")
	cookies, expiry, err := fetchCookies()
	if err != nil {
		logError("Can't fetch cookies: %v\n", err)
		return err
	}

	crumb, err := fetchCrumb(cookies)
	if err != nil {
		logError("Can't fetch crumb: %v\n", err)
		return err
	}

	client.crumb = crumb
	client.expiry = expiry
	client.cookies = cookies

	logDebug("Crumb refreshed: %s. Expires on %v\n", client.crumb, client.expiry)
	return nil
}

// newRequest creates a new Yahoo Finance request for the given path.
func (client *yClient) newRequest(path string) (*http.Request, error) {
	logInfo("Creating new request for path: %s\n", path)

	path = yFinURL + path
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		logError("Can't create api request: %v\n", err)
		return nil, err
	}

	req.Header = http.Header{
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Connection":      {"keep-alive"},
		"Content-Type":    {"application/json"},
		"Cookie":          {client.cookies},
		"Host":            {"query1.finance.yahoo.com"},
		"Origin":          {"https://finance.yahoo.com"},
		"Referer":         {"https://finance.yahoo.com"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-site"},
		"TE":              {"trailers"},
		"User-Agent":      {userAgent},
	}

	return req, nil
}

// do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshal the response
// into v. It also handles unmarshaling errors returned by the API.
func (client *yClient) do(req *http.Request, v interface{}) error {
	logInfo("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)

	start := time.Now()

	res, err := httpClient.Do(req)

	logDebug("Completed in %v\n", time.Since(start))

	if err != nil {
		logError("Request to api failed: %v\n", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logError("Can't close request response body: %v\n", err)
		}
	}(res.Body)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logError("Can't parse response: %v\n", err)
		return err
	}

	// convert the response to a string
	resBodyStr := string(resBody[:])
	logDebug("Response:\n%v\n", resBodyStr)

	// TODO Maybe add a way to access the raw json content for the users.
	if enableDump {
		// get the "symbol" value from the url
		symbol := req.URL.Query().Get("symbols")

		// format the body string to be indented
		var formattedBuffer bytes.Buffer
		json.Indent(&formattedBuffer, []byte(resBodyStr), "", "  ")

		// write the body string to a properly formatted json file named <symbol>_<MM>-<dd>_<hh>-<mm>.json located in the "c:\dump\Queries" folder
		// where symbol is the current symbol, and MM is the current month, dd the current day, hh the current local hour and mm the current local minute.
		// this is used to debug the api calls.
		//os.WriteFile(filepath.Join("c:\\dump\\Queries", symbol+"_"+time.Now().Format("01-02_15-04")+".json"), formattedBuffer.Bytes(), 0644)
		os.WriteFile(filepath.Join("c:\\dump\\Queries", symbol+".json"), formattedBuffer.Bytes(), 0644)
	}

	if res.StatusCode >= 400 {
		logError("API error: %q\n", resBody)
		return fmt.Errorf("error response received from upstream api: %s", res.Status)
	}

	logDebug("API response: %q\n", resBody)

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

// call is used by the public API methods to execute an API request.
func (client *yClient) call(path string, params queryParams, v interface{}) error {
	logInfo("Calling \"%s\" with params %v\n", path, params)

	// Check if the cookies have expired.
	if client.expiry.Before(time.Now()) {
		// Refresh the cookies and crumb.
		err := client.refreshCrumb()
		if err != nil {
			logError("Can't refresh crumb: %v\n", err)
			return err
		}
	}

	if client.crumb != "" {
		params["crumb"] = client.crumb
	}

	var values = url.Values{}
	if len(params) > 0 {
		for key, val := range params {
			values.Add(key, val)
		}
		path += "?" + values.Encode()
	}

	req, err := client.newRequest(path)
	if err != nil {
		logError("Can't create api request: %v\n", err)
		return err
	}

	return client.do(req, v)
}
