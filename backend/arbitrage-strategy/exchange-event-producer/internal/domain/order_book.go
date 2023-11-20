package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Order struct {
	Value       decimal.Decimal
	Quantity    decimal.Decimal
}

type OrderBook struct {
	Asks []Order
	Bids []Order
}

func (book OrderBook) String() string {
	return fmt.Sprintf("asks::%v bids::%v", book.Asks, book.Bids)
}