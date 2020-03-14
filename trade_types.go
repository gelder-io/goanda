package goanda

import (
	"encoding/json"
	"time"

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
	OpenTime              time.Time
	CloseTime             time.Time
	ClosingTransactionIDs []string
}

func (t *Trade) UnmarshalJSON(b []byte) (err error) {
	type _Trade Trade

	rawTrade := struct {
		_Trade

		OpenTime  string
		CloseTime string
	}{}

	if err = json.Unmarshal(b, &rawTrade); err != nil {
		return
	}

	*t = Trade(rawTrade._Trade)

	if t.OpenTime, err = time.Parse(time.RFC3339Nano, rawTrade.OpenTime); err != nil {
		return
	}

	if rawTrade.State == Closed {
		if t.CloseTime, err = time.Parse(time.RFC3339Nano, rawTrade.CloseTime); err != nil {
			return
		}
	}

	return
}

// TradeState represents the current state of a trade.
type TradeState string

const (
	Open               TradeState = "OPEN"
	Closed             TradeState = "CLOSED"
	CloseWhenTradeable TradeState = "CLOSE_WHEN_TRADEABLE"
)
