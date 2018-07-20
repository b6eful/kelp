package orderbook

import (
	"fmt"

	"github.com/lightyeario/kelp/model"
)

// OrderAction is the action of buy / sell
type OrderAction bool

// ActionBuy and ActionSell are the two actions
const (
	ActionBuy  OrderAction = false
	ActionSell OrderAction = true
)

// IsBuy returns true for buy actions
func (a OrderAction) IsBuy() bool {
	return a == ActionBuy
}

// IsSell returns true for sell actions
func (a OrderAction) IsSell() bool {
	return a == ActionSell
}

// String is the stringer function
func (a OrderAction) String() string {
	if a == ActionBuy {
		return "buy"
	} else if a == ActionSell {
		return "sell"
	}
	return "error, unrecognized order action"
}

var orderActionMap = map[string]OrderAction{
	"buy":  ActionBuy,
	"sell": ActionSell,
}

// OrderActionFromString is a convenience to convert from common strings to the corresponding OrderAction
func OrderActionFromString(s string) OrderAction {
	return orderActionMap[s]
}

// OrderType represents a type of an order, example market, limit, etc.
type OrderType int8

// These are the available order types
const (
	TypeMarket OrderType = 0
	TypeLimit  OrderType = 1
)

// IsMarket returns true for market orders
func (o OrderType) IsMarket() bool {
	return o == TypeMarket
}

// IsLimit returns true for limit orders
func (o OrderType) IsLimit() bool {
	return o == TypeLimit
}

// String is the stringer function
func (o OrderType) String() string {
	if o == TypeMarket {
		return "market"
	} else if o == TypeLimit {
		return "limit"
	}
	return "error, unrecognized order type"
}

var orderTypeMap = map[string]OrderType{
	"market": TypeMarket,
	"limit":  TypeLimit,
}

// OrderTypeFromString is a convenience to convert from common strings to the corresponding OrderType
func OrderTypeFromString(s string) OrderType {
	return orderTypeMap[s]
}

// Order represents an order in the orderbook
type Order struct {
	Pair        *model.TradingPair
	OrderAction OrderAction
	OrderType   OrderType
	Price       *model.Number
	Volume      *model.Number
	Timestamp   *model.Timestamp
}

// String is the stringer function
func (o Order) String() string {
	return fmt.Sprintf("Order[pair=%s, action=%s, type=%s, price=%s, vol=%s, ts=%d]",
		o.Pair,
		o.OrderAction,
		o.OrderType,
		o.Price.AsString(),
		o.Volume.AsString(),
		o.Timestamp.AsInt64(),
	)
}

// OrderBook encapsulates the concept of an orderbook on a market
type OrderBook struct {
	pair *model.TradingPair
	asks []Order
	bids []Order
}

// Asks returns the asks in an orderbook
func (o OrderBook) Asks() []Order {
	return o.asks
}

// Bids returns the bids in an orderbook
func (o OrderBook) Bids() []Order {
	return o.bids
}

// MakeOrderBook creates a new OrderBook from the asks and the bids
func MakeOrderBook(pair *model.TradingPair, asks []Order, bids []Order) *OrderBook {
	return &OrderBook{
		pair: pair,
		asks: asks,
		bids: bids,
	}
}

// TransactionID is typed for the concept of a transaction ID
type TransactionID string

// String is the stringer function
func (t TransactionID) String() string {
	return string(t)
}

// MakeTransactionID is a factory method for convenience
func MakeTransactionID(s string) *TransactionID {
	t := TransactionID(s)
	return &t
}

// OpenOrder represents an open order for a trading account
type OpenOrder struct {
	Order
	ID             string
	StartTime      *model.Timestamp
	ExpireTime     *model.Timestamp
	VolumeExecuted *model.Number
}

// String is the stringer function
func (o OpenOrder) String() string {
	return fmt.Sprintf("OpenOrder[order=%s, ID=%s, startTime=%d, expireTime=%d, volumeExecuted=%s]",
		o.Order.String(),
		o.ID,
		o.StartTime.AsInt64(),
		o.ExpireTime.AsInt64(),
		o.VolumeExecuted.AsString(),
	)
}
