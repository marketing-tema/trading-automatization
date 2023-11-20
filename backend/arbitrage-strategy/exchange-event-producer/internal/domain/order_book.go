package domain

import (
	"fmt"

	"algotrading/arbitrage-strategy/event-producer/common"
)

type OrderBook struct {
	Symbol string
	Asks   []common.PriceLevel
	Bids   []common.PriceLevel
}

func (book OrderBook) String() string {
	return fmt.Sprintf("LastUpdateId::%v asks::%v bids::%v", book.Symbol, book.Asks, book.Bids)
}
