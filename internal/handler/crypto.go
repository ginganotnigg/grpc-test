package handler

import (
	"context"
	"net/http"
	"encoding/json"
	"fmt"

	cryptopb "grpc-test/proto/crypto"
)


type CryptoServer struct {
	cryptopb.UnimplementedCryptoServiceServer
	coinGeckoURL string
}

func NewCryptoServer(coinGeckoURL string) *CryptoServer {
	return &CryptoServer{coinGeckoURL: coinGeckoURL}
}

func (s *CryptoServer) GetTopCoins(ctx context.Context, in *cryptopb.TopCoinsRequest) (*cryptopb.TopCoinsResponse, error) {
	resp, err := http.Get(s.coinGeckoURL)
	if err != nil {
			return nil, fmt.Errorf("failed to fetch data from CoinGecko: %v", err)
	}
	defer resp.Body.Close()

	var coins []cryptopb.Coin
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(coins) > 10 {
			coins = coins[:10]
	}

	coinPointers := make([]*cryptopb.Coin, len(coins))
	for i := range coins {
			coinPointers[i] = &coins[i]
	}

	return &cryptopb.TopCoinsResponse{Coins: coinPointers}, nil
}