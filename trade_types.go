package goanda

import (
	"github.com/shopspring/decimal"
)

// Trade represents a trade of an account.
type Trade struct {
	ID                    string
	Instrument            string
	State                 TradeState
	Price                 decimal.Decimal
	InitialUnits          decimal.Decimal
	InitialMarginRequired decimal.Decimal
	CurrentUnits          decimal.Decimal
	RealizedPL            decimal.Decimal
	UnrealizedPL          decimal.Decimal
	MarginUsed            decimal.Decimal
	AverageClosePrice     decimal.Decimal
	ClosingTransactionIDs []string
	OpenTime              string // TODO: should be time.Time, parsed according to RFC3339
	CloseTime             string // TODO: should be time.Time, parsed according to RFC3339
}

// TradeState represents the current state of a trade.
type TradeState string

const (
	Open               TradeState = "OPEN"
	Closed             TradeState = "CLOSED"
	CloseWhenTradeable TradeState = "CLOSE_WHEN_TRADEABLE"
)
