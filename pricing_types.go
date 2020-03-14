package goanda

import "github.com/shopspring/decimal"

// ClientPrice represents an account-specific price of an instrument.
type ClientPrice struct {
	Instrument string
	Time       string
	Tradeable  bool
	Bids       []PriceBucket
	Asks       []PriceBucket
}

// PriceBucket represents a price available for an amount of liquidity.
type PriceBucket struct {
	Price     decimal.Decimal
	Liquidity int
}
