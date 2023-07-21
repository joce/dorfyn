package dorfyn

import (
	"encoding/json"
	"fmt"
)

const (
	// apiErrorCode denotes an error caught by dorfyn stemming from invalid user inputs.
	apiErrorCode = "api-error"

	// remoteErrorCode denotes an error communicated in a response from a remote api source.
	remoteErrorCode = "remote-error"
)

// yError represents information returned in an error response from a Yahoo! finance call.
type yError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *yError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// CreateArgumentError returns an error with a message about missing arguments.
func CreateArgumentError(e string) error {
	return fmt.Errorf("code: %s, detail: %s", apiErrorCode, e)
}

// createRemoteError creates an error object from an error returned from a remote api source.
func createRemoteError(e error) error {
	return fmt.Errorf("code: %s, detail: %s", remoteErrorCode, e.Error())
}
