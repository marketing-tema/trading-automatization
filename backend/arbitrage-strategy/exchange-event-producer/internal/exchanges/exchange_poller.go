package exchanges

import "algotrading/arbitrage-strategy/event-producer/internal/domain"

type ExchangePoller interface {
	
	GetOrderBook(depth int, symbols []string) (*domain.OrderBook, error)
}