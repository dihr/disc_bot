package model

type TickerWrapper struct {
	// Wrappers Ticker data.
	Ticker Ticker `json:"ticker"`
}

type Ticker struct {
	// Highest trading unit price in the last 24 hours.
	High string `json:"high"`

	// Lowest unit trading price in the last 24 hours.
	Low string `json:"low"`

	// Amount traded in the last 24 hours.
	Vol string `json:"vol"`

	// Unit price of the last trade.
	Last string `json:"last"`

	// Highest bid price in the last 24 hours.
	Buy string `json:"buy"`

	// Lowest bid price in the last 24 hours.
	Sell string `json:"sell"`

	// Date and time information in Unix Era.
	Date int `json:"date"`
}
