package goanda

import "github.com/shopspring/decimal"

type orderRequest struct {
	Type         OrderType
	Instrument   string
	Units        decimal.Decimal
	TimeInForce  TimeInForce
	PositionFill OrderPositionFill
}

type MarketOrderRequest struct {
	orderRequest

	PriceBound decimal.Decimal
}

// TODO: Fill this with its fields.
type LimitOrderRequest struct {
	orderRequest
}

// TODO: Fill this with its fields.
type StopOrderRequest struct {
	orderRequest
}

// TODO: Fill this with its fields.
type MarketIfTouchedOrderRequest struct {
	orderRequest
}

// TODO: Fill this with its fields.
type TakeProfitOrderRequest struct {
	orderRequest
}

// TODO: Fill this with its fields.
type StopLossOrderRequest struct {
	orderRequest
}

// TODO: Fill this with its fields.
type TrailingStopLossOrderRequest struct {
	orderRequest
}

// OrderType represents the type of order to create.
type OrderType string

const (
	Market           OrderType = "MARKET"
	Limit            OrderType = "LIMIT"
	Stop             OrderType = "STOP"
	MarketIfTouched  OrderType = "MARKET_IF_TOUCHED"
	TakeProfit       OrderType = "TAKE_PROFIT"
	StopLoss         OrderType = "STOP_LOSS"
	TrailingStopLoss OrderType = "TRAILING_STOP_LOSS"
	FixedPrice       OrderType = "FIXED_PRICE"
)

// TimeInForce represents how long an order should remain pending before being automatically cancelled.
type TimeInForce string

const (
	GTC TimeInForce = "GTC"
	GTD TimeInForce = "GTD"
	GFD TimeInForce = "GFD"
	FOK TimeInForce = "FOK"
	IOC TimeInForce = "IOC"
)

// OrderPositionFill specifies how positions in the account are to be modified when an order is filled.
type OrderPositionFill string

const (
	OpenOnly    OrderPositionFill = "OPEN_ONLY"
	ReduceFirst OrderPositionFill = "REDUCE_FIRST"
	ReduceOnly  OrderPositionFill = "REDUCE_ONLY"
	Default     OrderPositionFill = "DEFAULT"
)

// isOrderRequest returns whether or not the given object is an order request.
func isOrderRequest(o interface{}) bool {
	switch o.(type) {
	case MarketOrderRequest:
		return true
	case LimitOrderRequest:
		return true
	case StopOrderRequest:
		return true
	case MarketIfTouchedOrderRequest:
		return true
	case TakeProfitOrderRequest:
		return true
	case StopLossOrderRequest:
		return true
	case TrailingStopLossOrderRequest:
		return true
	default:
		return false
	}
}
