package exchanges

import (
	"algotrading/arbitrage-strategy/event-producer/common"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type binancePollingClient struct {
	httpClient       *http.Client
	publicKey        string
	privateKey       string
	baseUrl          string
	supportedSymbols []string
}

type BinancePollingOptions struct {
	publicKey        string
	privateKey       string
	baseUrl          string
	supportedSymbols []string
}

func NewBinancePollingClient(options *BinancePollingOptions) (ExchangeRestPoller, error) {
	validation.ValidateStruct(options,
		validation.Field(options.publicKey, validation.Required),
		validation.Field(options.privateKey, validation.Required),
		validation.Field(options.supportedSymbols, validation.Required),
		validation.Field(options.baseUrl, validation.Required),
	)
	return &binancePollingClient{
		httpClient:       &http.Client{},
		supportedSymbols: options.supportedSymbols,
		baseUrl:          options.baseUrl,
		privateKey:       options.privateKey,
		publicKey:        options.publicKey,
	}, nil
}

func (client *binancePollingClient) GetOrderBook(ctx context.Context, depth int, symbol string) (*OrderBookResponse, error) {
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

	orderBookResponse := orderBookResponse{}
	err = dec.Decode(&orderBookResponse)

	if err != nil {
		return nil, err
	}

	return &OrderBookResponse{
		Bids: orderBookResponse.Bids,
		Asks: orderBookResponse.Asks,
	}, nil
}

type orderBookResponse struct {
	LastUpdateId int64               `json:"lastUpdateId"`
	Bids         []common.PriceLevel `json:"bids"`
	Asks         []common.PriceLevel `json:"asks"`
}
