package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"academic-certificate-verification-backend/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. Connect to local Hardhat blockchain
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	fmt.Println("Connected to Ethereum")

	// 2. Load Hardhat private key (NO 0x prefix)
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// 3. Create transaction signer
	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	fmt.Println("Using issuer address:", auth.From.Hex())

	// 4. Load DIDRegistry contract
	didRegistryAddress := common.HexToAddress(
		"0x5FbDB2315678afecb367f032d93F642f64180aa3",
	)

	didRegistry, err := blockchain.NewDIDRegistry(didRegistryAddress, client)
	if err != nil {
		log.Fatal("Failed to load DIDRegistry contract:", err)
	}

	// 5. Register issuer DID
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

	// 6. Wait until transaction is mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Failed while waiting for transaction:", err)
	}

	fmt.Println("Transaction mined in block:", receipt.BlockNumber)
}
