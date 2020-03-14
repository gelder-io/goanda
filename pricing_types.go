package goanda

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// ClientPrice represents an account-specific price of an instrument.
type ClientPrice struct {
	Instrument string
	Tradeable  bool
	Bids       []PriceBucket
	Asks       []PriceBucket
	Time       time.Time
}

func (p *ClientPrice) UnmarshalJSON(b []byte) (err error) {
	type _ClientPrice ClientPrice

	rawPrice := struct {
		_ClientPrice

		Time string
	}{}

	if err = json.Unmarshal(b, &rawPrice); err != nil {
		return
	}

	*p = ClientPrice(rawPrice._ClientPrice)

	if p.Time, err = time.Parse(time.RFC3339Nano, rawPrice.Time); err != nil {
		return
	}

	return
}

// PriceBucket represents a price available for an amount of liquidity.
type PriceBucket struct {
	Price     decimal.Decimal
	Liquidity int
}
