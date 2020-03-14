package goanda

import "github.com/shopspring/decimal"

// Position holds information about a market position.
type Position struct {
	Instrument   string
	PL           decimal.Decimal
	UnrealizedPL decimal.Decimal
	MarginUsed   decimal.Decimal
	ResettablePL decimal.Decimal
	Long         PositionSide
	Short        PositionSide
}

// PositionSide holds information either a long or short side of a position.
type PositionSide struct {
	TradeIDs     []string
	Units        decimal.Decimal
	AveragePrice decimal.Decimal
	PL           decimal.Decimal
	UnrealizedPL decimal.Decimal
	ResettablePL decimal.Decimal
}
