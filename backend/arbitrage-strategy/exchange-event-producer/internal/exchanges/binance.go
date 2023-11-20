package exchanges

import (
	"algotrading/arbitrage-strategy/event-producer/common"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type binanceClient struct {
	baseUrl          string
	supportedSymbols []string
	httpClient       *http.Client
}

type BinancePollingOptions struct {
	publicKey  string
	privateKey string
}

func NewBinancePollingClient(options *BinancePollingOptions) (ExchangeRestPoller, error) {
	return &binanceClient{}, nil
}

func (client *binanceClient) GetOrderBook(ctx context.Context, depth int, symbol string) (*OrderBookResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.baseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("depth", strconv.Itoa(depth))
	req.URL.RawQuery = q.Encode()
	res, err := client.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)

	obr := orderBookResponse{}
	err = dec.Decode(&obr)

	if err != nil {
		return nil, err
	}

	return &OrderBookResponse{
		Bids: obr.Bids,
		Asks: obr.Asks,
	}, nil
}

type orderBookResponse struct {
	LastUpdateId int64               `json:"lastUpdateId"`
	Bids         []common.PriceLevel `json:"bids"`
	Asks         []common.PriceLevel `json:"asks"`
}
