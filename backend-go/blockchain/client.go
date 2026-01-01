package blockchain

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	return client
}

func Context() context.Context {
	return context.Background()
}
