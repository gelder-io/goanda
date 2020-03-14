package goanda

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// Pricing represents an account-specific price of an instrument.
type Pricing struct {
	Instrument string
	Tradeable  bool
	Bids       []PriceBucket
	Asks       []PriceBucket
	Time       time.Time
}

func (p *Pricing) UnmarshalJSON(b []byte) (err error) {
	type _Pricing Pricing

	rawPrice := struct {
		_Pricing

		Time string
	}{}

	if err = json.Unmarshal(b, &rawPrice); err != nil {
		return
	}

	*p = Pricing(rawPrice._Pricing)

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
