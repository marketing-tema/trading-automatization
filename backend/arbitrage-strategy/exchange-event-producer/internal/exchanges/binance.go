package exchanges

import "algotrading/arbitrage-strategy/event-producer/internal/domain"

type binanceClient struct {
}

func NewBinanceClient(publicKey string, secretKey string) (ExchangePoller, error) {

	return &binanceClient{}, nil
}

func (client *binanceClient) GetOrderBook(depth int, symbols []string) (*domain.OrderBook, error) {
	return
}