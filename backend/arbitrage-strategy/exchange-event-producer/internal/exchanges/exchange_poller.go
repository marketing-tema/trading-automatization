package exchanges

import (
	"algotrading/arbitrage-strategy/event-producer/common"
	"context"
)

// uses poll strategy to request data from exchange
type ExchangeRestPoller interface {

	// requests order book for the given symbol of size
	GetOrderBook(ctx context.Context, depth int, symbol string) (*OrderBookResponse, error)
}

type OrderBookResponse struct {
	Bids []common.PriceLevel
	Asks []common.PriceLevel
}
