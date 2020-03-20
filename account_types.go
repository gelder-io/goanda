package goanda

import "github.com/shopspring/decimal"

// Account holds complete account information.
type Account struct {
	ID       string
	Alias    string
	Currency string

	OpenTradeCount    int
	OpenPositionCount int
	PendingOrderCount int

	HedgingEnabled bool

	MarginRate                 decimal.Decimal
	UnrealizedPL               decimal.Decimal
	NAV                        decimal.Decimal
	MarginUsed                 decimal.Decimal
	MarginAvailable            decimal.Decimal
	PositionValue              decimal.Decimal
	MarginCloseoutUnrealizedPL decimal.Decimal
	MarginCloseoutNAV          decimal.Decimal
	MarginCloseoutMarginUsed   decimal.Decimal
	MarginCloseoutPercent      decimal.Decimal
	WithdrawalLimit            decimal.Decimal
	Balance                    decimal.Decimal
	PL                         decimal.Decimal
	ResettablePL               decimal.Decimal

	Trades    []Trade
	Positions []Position
	// Orders    []Order
}

// AccountProperties holds properties related to an account.
type AccountProperties struct {
	ID   string
	Tags []string
}

// Instrument represents an instrument that can be traded by an account.
type Instrument struct {
	Name                        string
	DisplayName                 string
	Type                        InstrumentType
	PipLocation                 int
	DisplayPrecision            int
	TradeUnitsPrecision         int
	MinimumTradeSize            decimal.Decimal
	MaximumTrailingStopDistance decimal.Decimal
	MinimumTrailingStopDistance decimal.Decimal
	MaximumPositionSize         decimal.Decimal
	MaximumOrderUnits           decimal.Decimal
	MarginRate                  decimal.Decimal
	Commission                  InstrumentCommission
	Financing                   InstrumentFinancing
}

// InstrumentType represents what type of instrument it is.
type InstrumentType string

const (
	CurrencyType InstrumentType = "CURRENCY"
	CFDType      InstrumentType = "CFD"
	MetalType    InstrumentType = "METAL"
)

// InstrumentCommission represents instrument-specific commission information.
type InstrumentCommission struct {
	Commission        decimal.Decimal
	UnitsTraded       decimal.Decimal
	MinimumCommission decimal.Decimal
}

// InstrumentFinancing represents instrument-specific commission information.
type InstrumentFinancing struct {
	LongRate            decimal.Decimal
	ShortRate           decimal.Decimal
	FinancingDaysOfWeek []struct {
		DayOfWeek   DayOfWeek
		DaysCharged int
	}
}

type DayOfWeek string

const (
	Sunday    DayOfWeek = "SUNDAY"
	Monday    DayOfWeek = "MONDAY"
	Tuesday   DayOfWeek = "TUESDAY"
	Wednesday DayOfWeek = "WEDNESDAY"
	Thursday  DayOfWeek = "THURSDAY"
	Friday    DayOfWeek = "FRIDAY"
	Saturday  DayOfWeek = "SATURDAY"
)
