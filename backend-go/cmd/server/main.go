package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"academic-certificate-verification-backend/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Struct to read contracts.json
type ContractsConfig struct {
	DIDRegistry         string `json:"didRegistry"`
	CertificateRegistry string `json:"certificateRegistry"`
}

func main() {
	// 1. Connect to Hardhat node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	fmt.Println("Connected to Ethereum")

	// 2. Read deployed contract addresses
	configBytes, err := os.ReadFile("config/contracts.json")
	if err != nil {
		log.Fatal("Failed to read config/contracts.json:", err)
	}

	var contracts ContractsConfig
	if err := json.Unmarshal(configBytes, &contracts); err != nil {
		log.Fatal("Invalid contracts.json:", err)
	}

	didRegistryAddress := common.HexToAddress(contracts.DIDRegistry)

	// 3. Load private key (Hardhat default account #0)
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// 4. Create transaction signer
	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	fmt.Println("Using issuer address:", auth.From.Hex())

	// 5. Load DIDRegistry contract
	didRegistry, err := blockchain.NewDIDRegistry(didRegistryAddress, client)
	if err != nil {
		log.Fatal("Failed to load DIDRegistry contract:", err)
	}

	// 6. Sanity check: ensure contract exists
	code, err := client.CodeAt(context.Background(), didRegistryAddress, nil)
	if err != nil {
		log.Fatal("Failed to fetch contract code:", err)
	}
	if len(code) == 0 {
		log.Fatal("No contract code found at DIDRegistry address")
	}
	fmt.Println("Contract code found ✔")

	// 7. Register issuer DID
	tx, err := didRegistry.RegisterIssuer(
		auth,
		auth.From,
		"did:university:mits",
	)
	if err != nil {
		log.Fatal("Transaction failed:", err)
	}

	fmt.Println("Issuer registered successfully!")
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// 8. Wait for transaction confirmation
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Failed while waiting for transaction:", err)
	}

	fmt.Println("Transaction mined in block:", receipt.BlockNumber)
}
