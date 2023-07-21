package dorfyn

import (
	"github.com/shopspring/decimal"
)

// ChartBar is a single instance of a chart bar.
type ChartBar struct {
	Open      decimal.Decimal
	Low       decimal.Decimal
	High      decimal.Decimal
	Close     decimal.Decimal
	AdjClose  decimal.Decimal
	Volume    int
	Timestamp int
}

// OHLCHistoric is a historical quotation.
type OHLCHistoric struct {
	Open      float64
	Low       float64
	High      float64
	Close     float64
	AdjClose  float64
	Volume    int
	Timestamp int
}

// ChartMeta is metadata associated with a chart response.
type ChartMeta struct {
	Currency             string    `json:"currency" csv:"currency"`
	Symbol               string    `json:"symbol" csv:"symbol"`
	ExchangeName         string    `json:"exchangeName" csv:"exchangeName"`
	QuoteType            QuoteType `json:"instrumentType" csv:"instrumentType"`
	FirstTradeDate       int       `json:"firstTradeDate" csv:"firstTradeDate"`
	GMTOffset            int       `json:"GMTOffset" csv:"GMTOffset"`
	Timezone             string    `json:"timezone" csv:"timezone"`
	ExchangeTimezoneName string    `json:"exchangeTimezoneName" csv:"exchangeTimezoneName"`
	ChartPreviousClose   float64   `json:"chartPreviousClose" csv:"chartPreviousClose"`
	CurrentTradingPeriod struct {
		Pre struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			GMTOffset int    `json:"GMTOffset" csv:"GMTOffset"`
		} `json:"pre" csv:"pre_,inline"`
		Regular struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			GMTOffset int    `json:"GMTOffset" csv:"GMTOffset"`
		} `json:"regular" csv:"regular_,inline"`
		Post struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			GMTOffset int    `json:"GMTOffset" csv:"GMTOffset"`
		} `json:"post" csv:"post_,inline"`
	} `json:"currentTradingPeriod" csv:"currentTradingPeriod_,inline"`
	DataGranularity string   `json:"dataGranularity" csv:"dataGranularity"`
	ValidRanges     []string `json:"validRanges" csv:"-"`
}

// OptionsMeta is metadata associated with an options' response.
type OptionsMeta struct {
	UnderlyingSymbol   string    `json:"underlyingSymbol" csv:"underlyingSymbol"`
	ExpirationDate     int       `json:"expirationDate" csv:"expirationDate"`
	AllExpirationDates []int     `json:"allExpirationDates" csv:"-"`
	Strikes            []float64 `json:"strikes" csv:"-"`
	HasMiniOptions     bool      `json:"hasMiniOptions"`
	OldQuote           *OldQuote `json:"quote,omitempty" csv:"quote_,inline"`
}

// Straddle is a put/call straddle for a particular strike.
type Straddle struct {
	Strike float64   `json:"strike" csv:"strike"`
	Call   *Contract `json:"call,omitempty" csv:"call_,inline"`
	Put    *Contract `json:"put,omitempty" csv:"put_,inline"`
}

// Contract is a struct containing a single option contract, usually part of a chain.
type Contract struct {
	Symbol            string  `json:"contractSymbol" csv:"contractSymbol"`
	Strike            float64 `json:"strike" csv:"strike"`
	Currency          string  `json:"currency" csv:"currency"`
	LastPrice         float64 `json:"lastPrice" csv:"lastPrice"`
	Change            float64 `json:"change" csv:"change"`
	PercentChange     float64 `json:"percentChange" csv:"percentChange"`
	Volume            int     `json:"volume" csv:"volume"`
	OpenInterest      int     `json:"openInterest" csv:"openInterest"`
	Bid               float64 `json:"bid" csv:"bid"`
	Ask               float64 `json:"ask" csv:"ask"`
	Size              string  `json:"contractSize" csv:"contractSize"`
	Expiration        int     `json:"expiration" csv:"expiration"`
	LastTradeDate     int     `json:"lastTradeDate" csv:"lastTradeDate"`
	ImpliedVolatility float64 `json:"impliedVolatility" csv:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney" csv:"inTheMoney"`
}
