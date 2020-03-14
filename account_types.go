package goanda

import "github.com/shopspring/decimal"

// Account holds complete account information.
type Account struct {
	ID                         string
	Alias                      string
	Currency                   string
	OpenTradeCount             int
	OpenPositionCount          int
	PendingOrderCount          int
	HedgingEnabled             bool
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
	Trades                     []Trade
	Positions                  []Position
	// Orders                     []Order
}

// AccountProperties holds properties related to an account.
type AccountProperties struct {
	ID   string
	Tags []string
}
