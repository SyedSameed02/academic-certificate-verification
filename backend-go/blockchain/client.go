package blockchain

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Eth      *ethclient.Client
	Auth     *bind.TransactOpts
	CallOpts *bind.CallOpts
}

func NewClient() *Client {
	rpc := os.Getenv("RPC_URL")
	key := os.Getenv("PRIVATE_KEY")

	eth, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}



	key = strings.TrimPrefix(key, "0x")
	if len(key) != 64 {
		log.Fatalf("PRIVATE_KEY must be 64 hex chars (32 bytes), got %d", len(key))
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		bigChainID(),
	)
	if err != nil {
		log.Fatal(err)
	}

	auth.From = crypto.PubkeyToAddress(*publicKey)

	return &Client{
		Eth:  eth,
		Auth: auth,
		CallOpts: &bind.CallOpts{
			Context: context.Background(),
		},
	}
}
func bigChainID() *big.Int {
    id := os.Getenv("CHAIN_ID")
    chainID, _ := new(big.Int).SetString(id, 10)
    return chainID
}
