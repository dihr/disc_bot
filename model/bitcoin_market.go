package model

type TickerWrapper struct {
	// Wrappers Ticker data.
	Ticker Ticker `json:"ticker"`
}

type DaySummary struct {
	// Daily summary date.
	Date string `json:"date"`

	// Unit price for opening trading on the day.
	Opening float64 `json:"opening"`

	// Unit price at the close of trading on the day.
	Closing float64 `json:"closing"`

	// Lowest trading unit price on the day.
	Lowest float64 `json:"lowest"`

	// Highest unit trading price on the day.
	Highest float64 `json:"highest"`

	// Volume of Reais (BRL) traded on the day.
	Volume string `json:"volume"`

	// Amount of digital currency traded on the day.
	Quantity string `json:"quantity"`

	// Number of trades carried out on the day.
	Amount int `json:"amount"`

	// Average unit price of trading on the day.
	AvgPrice float64 `json:"avg_price"`
}

type OrderBook struct {
	// List of sales offers, ordered from lowest to highest price.
	Asks [][]float64 `json:"asks"`

	// List of shopping offers, ordered from highest to lowest price.
	Bids [][]float64 `json:"bids"`
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
