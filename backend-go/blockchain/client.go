package blockchain

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Shared instances
var (
	EthClient *ethclient.Client
	initOnce  sync.Once
)

// Init initializes ethereum client once
func Init(rpcURL string) error {
	var err error

	initOnce.Do(func() {
		EthClient, err = ethclient.Dial(rpcURL)
		if err != nil {
			log.Println("❌ Failed to connect to Ethereum:", err)
			return
		}
		log.Println("✅ Ethereum client connected")
	})

	return err
}
