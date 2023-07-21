package dorfyn

import "strings"

// quoteResponse is a yfin quote quoteResponse.
type quoteResponse struct {
	Inner struct {
		Result []Quote `json:"result"`
		Error  *yError `json:"error"`
	} `json:"quoteResponse"`
}

const (
	// yFinQuoteAPI is the path to the Yahoo! finance quote API.
	yFinQuoteAPI string = "/v7/finance/quote"
)

var (
	client yClient
)

// GetQuotes returns quotes for the given symbols.
func GetQuotes(symbols []string) ([]Quote, error) {
	if len(symbols) == 0 {
		return nil, CreateArgumentError("No symbols provided to GetQuotes")
	}

	params := map[string]string{"symbols": strings.Join(symbols, ",")}
	resp := quoteResponse{}

	err := client.call(yFinQuoteAPI, params, &resp)
	if err != nil {
		return nil, createRemoteError(err)
	}

	if resp.Inner.Error != nil {
		err = createRemoteError(resp.Inner.Error)
	}

	return resp.Inner.Result, err
}

type (
	// QuoteType alias for asset classification.
	QuoteType string
	// MarketState alias for market state.
	MarketState string
	// OptionType alias for option type.
	OptionType string
)

const (
	// QuoteTypeEquity the returned quote for an equity.
	QuoteTypeEquity QuoteType = "EQUITY"
	// QuoteTypeIndex the returned quote for an index.
	QuoteTypeIndex QuoteType = "INDEX"
	// QuoteTypeOption the returned quote for an option contract.
	QuoteTypeOption QuoteType = "OPTION"
	// QuoteTypeForexPair the returned quote for a forex pair.
	QuoteTypeForexPair QuoteType = "CURRENCY"
	// QuoteTypeCryptoPair the returned quote for a crypto pair.
	QuoteTypeCryptoPair QuoteType = "CRYPTOCURRENCY"
	// QuoteTypeFuture the returned quote for a futures contract.
	QuoteTypeFuture QuoteType = "FUTURE"
	// QuoteTypeETF the returned quote for an etf.
	QuoteTypeETF QuoteType = "ETF"
	// QuoteTypeMutualFund the returned quote for a mutual fund.
	QuoteTypeMutualFund QuoteType = "MUTUALFUND"

	// MarketStatePrePre pre-pre market state.
	MarketStatePrePre MarketState = "PREPRE"
	// MarketStatePre pre market state. Usually weekdays from 4:00am - 9:30am Eastern, excluding holidays.
	MarketStatePre MarketState = "PRE"
	// MarketStateRegular regular market state. Usually weekdays from 9:30am - 4:00pm Eastern, excluding holidays.
	MarketStateRegular MarketState = "REGULAR"
	// MarketStatePost post market state. Usually weekdays from 4:00pm - 8:00pm Eastern, excluding holidays.
	MarketStatePost MarketState = "POST"
	// MarketStatePostPost post-post market state.
	MarketStatePostPost MarketState = "POSTPOST"
	// MarketStateClosed closed market state. Usually weekdays from 8:00pm - 4:00am Eastern, weekends from 8:00pm Friday to 4:00am Monday Eastern and certain holidays.
	MarketStateClosed MarketState = "CLOSED"

	// OptionTypeCall call option type.
	OptionTypeCall OptionType = "CALL"
	// OptionTypePut put option type.
	OptionTypePut OptionType = "PUT"
)

// Quote represents a quote for a stock or security.
type Quote struct {
	// Ask is the asking price, or the lowest price that a seller is willing to accept for a unit of the security. Applies to CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	Ask *float64 `json:"ask,omitempty"`
	// AskSize is the total number of shares that are currently being asked for at the asking price. Applies to CURRENCY, EQUITY, ETF and INDEX quotes.
	AskSize *float64 `json:"askSize,omitempty"`
	// AverageAnalystRating is a measure of the consensus recommendation for a given stock by financial analysts. Applies to EQUITY quotes.
	AverageAnalystRating *string `json:"averageAnalystRating,omitempty"`
	// AverageDailyVolume10Day is the average number of shares traded each day over the last 10 days. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	AverageDailyVolume10Day *int `json:"averageDailyVolume10Day,omitempty"`
	// AverageDailyVolume3Month is the average number of shares traded each day over the last 3 months. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	AverageDailyVolume3Month *int `json:"averageDailyVolume3Month,omitempty"`
	// Bid is the bid price, or the highest price that a buyer is willing to pay for a unit of the security. Applies to CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	Bid *float64 `json:"bid,omitempty"`
	// BidSize is the total number of shares that buyers want to buy at the bid price. Applies to CURRENCY, EQUITY, ETF and INDEX quotes.
	BidSize *int `json:"bidSize,omitempty"`
	// BookValue is the net asset value of a company, calculated by total assets minus intangible assets (patents, goodwill) and liabilities. Applies to EQUITY, ETF and MUTUALFUND quotes.
	BookValue *float64 `json:"bookValue,omitempty"`
	// CirculatingSupply, in the context of cryptocurrencies, is the amount of coins that are publicly available and circulating in the market. Applies to CRYPTOCURRENCY quotes.
	CirculatingSupply *int `json:"circulatingSupply,omitempty"`
	// CoinImageUrl is the URL of the image representing the cryptocurrency. Applies to CRYPTOCURRENCY quotes.
	CoinImageUrl *string `json:"coinImageUrl,omitempty"`
	// CoinMarketCapLink is the URL of the MarketCap site for the cryptocurrency. Applies to CRYPTOCURRENCY quotes.
	CoinMarketCapLink *string `json:"coinMarketCapLink,omitempty"`
	// ContractSymbol represents the ticker symbol for a futures contract. Applies to FUTURE quotes.
	ContractSymbol *bool `json:"contractSymbol,omitempty"`
	// CryptoTradeable is a boolean value indicating whether the cryptocurrency can be traded. Applies to CRYPTOCURRENCY quotes.
	CryptoTradeable *bool `json:"cryptoTradeable,omitempty"`
	// Currency is the currency in which the security is traded. Applies to ALL quotes.
	Currency *string `json:"currency"`
	// CustomPriceAlertConfidence represents a value whose meaning is not clear at the moment. Seen values have been NONE, LOW and HIGH. Applies to ALL quotes.
	CustomPriceAlertConfidence *string `json:"customPriceAlertConfidence"`
	// DisplayName is the user-friendly name of the stock or security. Applies to EQUITY quotes.
	DisplayName *string `json:"displayName,omitempty"`
	// DividendDate is the date when the company is expected to pay its next dividend. Applies to EQUITY, ETF and MUTUALFUND quotes.
	DividendDate *int `json:"dividendDate,omitempty"`
	// DividendRate is the amount of dividends that a company is expected to pay over the next year. Applies to MUTUALFUND quotes.
	DividendRate *float64 `json:"dividendRate,omitempty"`
	// DividendYield is a financial ratio that indicates how much a company pays out in dividends each year relative to its stock price. Applies to ETF and MUTUALFUND quotes.
	DividendYield *float64 `json:"dividendYield,omitempty"`
	// EarningsTimestamp is the timestamp of the company's earnings announcement. Applies to EQUITY quotes.
	EarningsTimestamp *int `json:"earningsTimestamp,omitempty"`
	// EpsCurrentYear is the company's earnings per share (EPS) for the current year. Applies to EQUITY quotes.
	EarningsTimestampEnd *int `json:"earningsTimestampEnd,omitempty"`
	// EpsForward is the company's projected earnings per share (EPS) for the next fiscal year. Applies to EQUITY quotes.
	EarningsTimestampStart *int `json:"earningsTimestampStart,omitempty"`
	// EpsTrailingTwelveMonths is the company's earnings per share (EPS) for the past 12 months. Applies to EQUITY, ETF and MUTUALFUND quotes.
	EpsCurrentYear *float64 `json:"epsCurrentYear,omitempty"`
	// EpsForward is the company's projected earnings per share (EPS) for the next fiscal year. Applies to EQUITY quotes.
	EpsForward *float64 `json:"epsForward,omitempty"`
	// EpsTrailingTwelveMonths is the company's earnings per share (EPS) for the past 12 months. Applies to EQUITY, ETF and MUTUALFUND quotes.
	EpsTrailingTwelveMonths *float64 `json:"epsTrailingTwelveMonths,omitempty"`
	// EsgPopulated is a boolean indicating whether the company's environmental, social, and governance (ESG) ratings are populated. Applies to ALL quotes
	EsgPopulated *bool `json:"esgPopulated"`
	// Exchange is the securities exchange on which the security is traded. Applies to ALL quotes.
	Exchange *string `json:"exchange"`
	// ExchangeDataDelayedBy is the delay in data from the exchange, typically in minutes. Applies to ALL quotes.
	ExchangeDataDelayedBy *int `json:"exchangeDataDelayedBy"`
	// ExchangeTimezoneName is the name of the timezone of the exchange. Applies to ALL quotes.
	ExchangeTimezoneName *string `json:"exchangeTimezoneName"`
	// ExchangeTimezoneShortName is the short name of the timezone of the exchange. Applies to ALL quotes.
	ExchangeTimezoneShortName *string `json:"exchangeTimezoneShortName"`
	// ExpireDate is the date on which the option contract expires. Applies to OPTION quotes.
	ExpireDate *int `json:"expireDate,omitempty"`
	// ExpireIsoDate is the date on which the option contract expires, in ISO 8601 format. Applies to OPTION quotes.
	ExpireIsoDate *string `json:"expireIsoDate,omitempty"`
	// FiftyDayAverage is the average closing price of the stock over the past 50 trading days. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	FiftyDayAverage *float64 `json:"fiftyDayAverage,omitempty"`
	// FiftyDayAverageChange is the change in the 50-day average price from the previous trading day. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	FiftyDayAverageChange *float64 `json:"fiftyDayAverageChange,omitempty"`
	// FiftyDayAverageChangePercent is the percent change in the 50-day average price from the previous trading day. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	FiftyDayAverageChangePercent *float64 `json:"fiftyDayAverageChangePercent,omitempty"`
	// FiftyTwoWeekChangePercent is the percentage change in price over the past 52 weeks. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	FiftyTwoWeekChangePercent *float64 `json:"fiftyTwoWeekChangePercent,omitempty"`
	// FiftyTwoWeekHigh is the highest price the stock has traded at in the past 52 weeks. Applies to ALL quotes.
	FiftyTwoWeekHigh *float64 `json:"fiftyTwoWeekHigh"`
	// FiftyTwoWeekHighChange is the change in the 52-week high price from the previous trading day. Applies to ALL quotes.
	FiftyTwoWeekHighChange *float64 `json:"fiftyTwoWeekHighChange"`
	// FiftyTwoWeekHighChangePercent is the percent change in the 52-week high price from the previous trading day. Applies to ALL quotes.
	FiftyTwoWeekHighChangePercent *float64 `json:"fiftyTwoWeekHighChangePercent"`
	// FiftyTwoWeekLow is the lowest price the stock has traded at in the past 52 weeks. Applies to ALL quotes.
	FiftyTwoWeekLow *float64 `json:"fiftyTwoWeekLow"`
	// FiftyTwoWeekLowChange is the change in the 52-week low price from the previous trading day. Applies to ALL quotes.
	FiftyTwoWeekLowChange *float64 `json:"fiftyTwoWeekLowChange"`
	// FiftyTwoWeekLowChangePercent is the percent change in the 52-week low price from the previous trading day. Applies to ALL quotes.
	FiftyTwoWeekLowChangePercent *float64 `json:"fiftyTwoWeekLowChangePercent"`
	// FiftyTwoWeekRange is the range of the highest and lowest prices the stock has traded at over the past 52 weeks. Applies to ALL quotes.
	FiftyTwoWeekRange *string `json:"fiftyTwoWeekRange"`
	// FinancialCurrency is the currency in which the company reports its financial results. Applies to EQUITY, ETF and MUTUALFUND quotes.
	FinancialCurrency *string `json:"financialCurrency,omitempty"`
	// FirstTradeDateMilliseconds is the timestamp of the first trade of this security, in milliseconds. Applies to ALL quotes.
	FirstTradeDateMilliseconds *int `json:"firstTradeDateMilliseconds"`
	// ForwardPE is the forward price-to-earnings ratio, calculated as the current share price divided by projected earnings per share for the next 12 months. Applies to EQUITY quotes.
	ForwardPE *float64 `json:"forwardPE,omitempty"`
	// FromCurrency is, in a currency pair, the currency that is being exchanged from. Applies to CRYPTOCURRENCY quotes.
	FromCurrency *string `json:"fromCurrency,omitempty"`
	// FullExchangeName is the full name of the securities exchange on which the security is traded. Applies to ALL quotes.
	FullExchangeName *string `json:"fullExchangeName"`
	// GmtOffSetMilliseconds is the offset from GMT of the exchange, in milliseconds. Applies to ALL quotes.
	GmtOffSetMilliseconds *int `json:"gmtOffSetMilliseconds"`
	// HeadSymbolAsString is the symbol of the contract's underlying security. Applies to OPTION quotes.
	HeadSymbolAsString *string `json:"headSymbolAsString,omitempty"`
	// IpoExpectedDate is the expected date of the initial public offering (IPO). Applies to EQUITY quotes.
	IpoExpectedDate *string `json:"ipoExpectedDate,omitempty"`
	// Language is the language in which financial results are reported. Applies to ALL quotes.
	Language *string `json:"language"`
	// LastMarket is the last market in which the security was traded. Applies to CRYPTOCURRENCY quotes.
	LastMarket *string `json:"lastMarket,omitempty"`
	// LogoUrl is the URL of the coin logo. Applies to CRYPTOCURRENCY quotes.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// LongName is the official name of the company. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, INDEX and MUTUALFUND quotes.
	LongName *string `json:"longName,omitempty"`
	// Market is the market in which the security is primarily traded. Applies to ALL quotes.
	Market *string `json:"market"`
	// MarketCap is the market capitalization of the company, calculated as share price times the number of outstanding shares. Applies to CRYPTOCURRENCY, EQUITY, ETF and MUTUALFUND quotes.
	MarketCap *int `json:"marketCap,omitempty"`
	// MarketState represents the current state of the market for a security.
	MarketState *MarketState `json:"marketState,omitempty"`
	// MessageBoardId is the identifier for the Yahoo! Finance message board for this security. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, INDEX and MUTUALFUND quotes.
	MessageBoardId *string `json:"messageBoardId,omitempty"`
	// NameChangeDate is the date on which the company last changed its name. Applies to EQUITY quotes.
	NameChangeDate *string `json:"nameChangeDate,omitempty"`
	// NetAssets is the total net assets of the company. Applies to ETF and MUTUALFUND quotes.
	NetAssets *float64 `json:"netAssets,omitempty"`
	// NetExpenseRatio is the ratio of total expenses to total net assets. Applies to ETF and MUTUALFUND quotes.
	NetExpenseRatio *float64 `json:"netExpenseRatio,omitempty"`
	// OpenInterest is the total number of open contracts on a futures or options market. Applies to FUTURE and OPTION quotes.
	OpenInterest *int `json:"openInterest,omitempty"`
	// OptionType is the type of option. Applies to OPTION quotes.
	OptionsType *OptionType `json:"optionsType,omitempty"`
	// PostMarketChange is the change in the security's price in post-market trading. Applies to ALL quotes.
	PostMarketChange *float64 `json:"postMarketChange"`
	// PostMarketChangePercent is the percent change in the security's price in post-market trading. Applies to ALL quotes.
	PostMarketChangePercent *float64 `json:"postMarketChangePercent"`
	// PostMarketPrice is the price of the security in post-market trading. Applies to ALL quotes.
	PostMarketPrice *float64 `json:"postMarketPrice"`
	// PostMarketTime is the time of the most recent post-market trade. Applies to ALL quotes.
	PostMarketTime *int `json:"postMarketTime"`
	// PreMarketChange is the change in the security's price in pre-market trading. Applies to ALL quotes.
	PreMarketChange *float64 `json:"preMarketChange"`
	// PreMarketChangePercent is the percent change in the security's price in pre-market trading. Applies to ALL quotes.
	PreMarketChangePercent *float64 `json:"preMarketChangePercent"`
	// PreMarketPrice is the price of the security in pre-market trading. Applies to ALL quotes.
	PreMarketPrice *float64 `json:"preMarketPrice"`
	// PreMarketTime is the time of the most recent pre-market trade. Applies to ALL quotes.
	PreMarketTime *int `json:"preMarketTime"`
	// PrevName is the name of the company prior to its most recent name change. Applies to EQUITY quotes.
	PrevName *string `json:"prevName,omitempty"`
	// PriceEpsCurrentYear is the price of the stock divided by the company's earnings per share (EPS) for the current year. Applies to EQUITY quotes.
	PriceEpsCurrentYear *float64 `json:"priceEpsCurrentYear,omitempty"`
	// PriceHnt is a hint about the precision of the price data (e.g. the number of decimal places). Applies to ALL quotes.
	PriceHint *int `json:"priceHint"`
	// PriceToBook is the price-to-book ratio, calculated as the market price per share divided by the book value per share. Applies to EQUITY, ETF and MUTUALFUND quotes.
	PriceToBook *float64 `json:"priceToBook,omitempty"`
	// QuoteSourceName is the name of the source providing the quote. Applies to ALL quotes.
	QuoteSourceName *string `json:"quoteSourceName"`
	// QuoteType is the type of quote. Applies to ALL quotes.
	QuoteType *QuoteType `json:"quoteType"`
	// Region is the region in which the company is located. Applies to ALL quotes.
	Region *string `json:"region"`
	// RegularMarketChange is the change in the security's price from the previous regular trading session. Applies to ALL quotes.
	RegularMarketChange *float64 `json:"regularMarketChange"`
	// RegularMarketChangePercent is the percent change in the security's price from the previous regular trading session. Applies to ALL quotes.
	RegularMarketChangePercent *float64 `json:"regularMarketChangePercent"`
	// RegularMarketDayHigh is the highest price at which the security has traded during the most recent regular trading session. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	RegularMarketDayHigh *float64 `json:"regularMarketDayHigh,omitempty"`
	// RegularMarketDayLow is the lowest price at which the security has traded during the most recent regular trading session. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	RegularMarketDayLow *float64 `json:"regularMarketDayLow,omitempty"`
	// RegularMarketDayRange is the range of prices at which the security has traded during the most recent regular trading session. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	RegularMarketDayRange *string `json:"regularMarketDayRange,omitempty"`
	// RegularMarketOpen is the price at which the security first traded in the most recent regular trading session. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	RegularMarketOpen *float64 `json:"regularMarketOpen,omitempty"`
	// RegularMarketPreviousClose is the closing price of the security in the previous regular trading session. Applies to ALL quotes.
	RegularMarketPreviousClose *float64 `json:"regularMarketPreviousClose"`
	// RegularMarketPrice is the last traded price of the security in the most recent regular trading session. Applies to ALL quotes.
	RegularMarketPrice *float64 `json:"regularMarketPrice"`
	// RegularMarketTime is the time of the most recent trade in the regular trading session. Applies to ALL quotes.
	RegularMarketTime *int `json:"regularMarketTime"`
	// RegularMarketVolume is the number of shares traded during the most recent regular trading session. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and OPTION quotes.
	RegularMarketVolume *int `json:"regularMarketVolume,omitempty"`
	// SharesOutstanding is the number of shares currently held by all shareholders. Applies to EQUITY, ETF and MUTUALFUND quotes.
	SharesOutstanding *int `json:"sharesOutstanding,omitempty"`
	// ShortName is a short, user-friendly name for the stock or security. Applies to ALL quotes.
	ShortName *string `json:"shortName"`
	// SourceInterval is the interval at which the data source provides updates, in seconds. Applies to ALL quotes.
	SourceInterval *int `json:"sourceInterval"`
	// StartDate is the date on which the coin started trading. Applies to CRYPTOCURRENCY.
	StartDate *int `json:"startDate,omitempty"`
	// Strike is the strike price of an options contract, which is the price at which the contract can be exercised. Applies to OPTION quotes.
	Strike *float64 `json:"strike,omitempty"`
	// Symbol is the ticker symbol of the security. Applies to ALL quotes.
	Symbol *string `json:"symbol"`
	// ToCurrency is, in a currency pair, is the currency that is being exchanged to. Applies to CRYPOTOCURRENCY quotes.
	ToCurrency *string `json:"toCurrency,omitempty"`
	// Tradeable is a boolean value indicating whether the security is currently tradeable. Applies to ALL quotes.
	Tradeable *bool `json:"tradeable"`
	// TrailingAnnualDividendRate is the company's dividend payment per share over the past 12 months. Applies to EQUITY, ETF and MUTUALFUND quotes.
	TrailingAnnualDividendRate *float64 `json:"trailingAnnualDividendRate,omitempty"`
	// TrailingAnnualDividendYield is the annual dividend payment divided by the current stock price. Applies to EQUITY, ETF and MUTUALFUND quotes.
	TrailingAnnualDividendYield *float64 `json:"trailingAnnualDividendYield,omitempty"`
	// TrailingPE is the trailing price-to-earnings ratio, calculated as the current share price divided by the earnings per share (EPS) over the past 12 months. Applies to EQUITY, ETF and MUTUALFUND quotes.
	TrailingPE *float64 `json:"trailingPE,omitempty"`
	// TrailingThreeMonthNavReturns are the trailing 3-month net asset value (NAV) returns. Applies to ETF quotes.
	TrailingThreeMonthNavReturns *float64 `json:"trailingThreeMonthNavReturns,omitempty"`
	// TrailingThreeMonthReturns are the trailing 3-month returns for the security. Applies to ETF and MUTUALFUND quotes.
	TrailingThreeMonthReturns *float64 `json:"trailingThreeMonthReturns,omitempty"`
	// Triggerable represents a Boolean value whose meaning is not clear at the moment. Applies to ALL quotes.
	Triggerable *bool `json:"triggerable"`
	// TwoHundredDayAverage is the average closing price of the stock over the past 200 trading days. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	TwoHundredDayAverage *float64 `json:"twoHundredDayAverage,omitempty"`
	// TwoHundredDayAverageChange is the change in the 200-day average price from the previous trading day. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	TwoHundredDayAverageChange *float64 `json:"twoHundredDayAverageChange,omitempty"`
	// TwoHundredDayAverageChangePercent is the percent change in the 200-day average price from the previous trading day. Applies to CRYPTOCURRENCY, CURRENCY, EQUITY, ETF, FUTURE, INDEX and MUTUALFUND quotes.
	TwoHundredDayAverageChangePercent *float64 `json:"twoHundredDayAverageChangePercent,omitempty"`
	// TypeDisp is a user-friendly representation of the QuoteType. Applies to ALL quotes.
	TypeDisp *string `json:"typeDisp"`
	// UnderlyingExchangeSymbol is the symbol of the exchange on which the underlying security of a derivative is traded. Applies to FUTURE quotes.
	UnderlyingExchangeSymbol *string `json:"underlyingExchangeSymbol,omitempty"`
	// UnderlyingShortName is the short name of the underlying security of a derivative. Applies to OPTION quotes.
	UnderlyingShortName *string `json:"underlyingShortName,omitempty"`
	// UnderlyingSymbol is the ticker symbol of the underlying security of a derivative. Applies to FUTURE and OPTION quotes.
	UnderlyingSymbol *string `json:"underlyingSymbol,omitempty"`
	// Volume24Hr is the total trading volume of a cryptocurrency in the past 24 hours. Applies to CRYPTOCURRENCY quotes.
	Volume24Hr *int `json:"volume24Hr,omitempty"`
	// VolumeAllCurrencies is the total trading volume of a cryptocurrency across all currencies in the past 24 hours. Applies to CRYPTOCURRENCY quotes.
	VolumeAllCurrencies *int `json:"volumeAllCurrencies,omitempty"`
	// YtdReturn is the year-to-date return on the security. Applies to ETF and MUTUALFUND quotes.
	YtdReturn *float64 `json:"ytdReturn,omitempty"`
}
