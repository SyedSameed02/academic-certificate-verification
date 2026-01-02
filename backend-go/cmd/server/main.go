package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"academic-certificate-verification-backend/blockchain"
	"academic-certificate-verification-backend/internal/certificate"

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
	ctx := context.Background()

	// 1️⃣ Connect to Hardhat node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	fmt.Println("✅ Connected to Ethereum")

	// 2️⃣ Read deployed contract addresses
	configBytes, err := os.ReadFile("config/contracts.json")
	if err != nil {
		log.Fatal("Failed to read config/contracts.json:", err)
	}

	var contracts ContractsConfig
	if err := json.Unmarshal(configBytes, &contracts); err != nil {
		log.Fatal("Invalid contracts.json:", err)
	}

	didRegistryAddress := common.HexToAddress(contracts.DIDRegistry)
	certificateRegistryAddress := common.HexToAddress(contracts.CertificateRegistry)

	// 3️⃣ Load private key (Hardhat default account #0)
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// 4️⃣ Create transaction signer
	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}
	auth.GasLimit = 3_000_000

	fmt.Println("🧾 Using issuer address:", auth.From.Hex())

	// 5️⃣ Load DIDRegistry contract
	didRegistry, err := blockchain.NewDIDRegistry(didRegistryAddress, client)
	if err != nil {
		log.Fatal("Failed to load DIDRegistry:", err)
	}

	// 6️⃣ Sanity check: ensure contract exists
	code, err := client.CodeAt(ctx, didRegistryAddress, nil)
	if err != nil {
		log.Fatal("Failed to fetch contract code:", err)
	}
	if len(code) == 0 {
		log.Fatal("No contract code found at DIDRegistry address")
	}
	fmt.Println("✅ DIDRegistry contract found")

	// 7️⃣ Register issuer DID
	fmt.Println("🚀 Registering issuer...")
	registerTx, err := didRegistry.RegisterIssuer(
		auth,
		auth.From,
		"did:university:mits",
	)
	if err != nil {
		log.Fatal("Issuer registration failed:", err)
	}

	_, err = bind.WaitMined(ctx, client, registerTx)
	if err != nil {
		log.Fatal("Failed waiting for issuer registration tx:", err)
	}
	fmt.Println("✅ Issuer registered")

	// 8️⃣ Create certificate data
	cert := certificate.CertificateData{
		StudentName: "Rahul Sharma",
		RollNumber:  "CSE2021012",
		Course:      "B.Tech Computer Science",
		University:  "MITS",
		Year:        "2025",
	}

	// 9️⃣ Hash certificate
	certHash := certificate.HashCertificate(cert)
	fmt.Println("🔐 Certificate hash:", certHash)

	// 🔟 Load CertificateRegistry contract
	certificateRegistry, err := blockchain.NewCertificateRegistry(
		certificateRegistryAddress,
		client,
	)
	if err != nil {
		log.Fatal("Failed to load CertificateRegistry:", err)
	}

	// 1️⃣1️⃣ Issue certificate
	fmt.Println("🚀 Issuing certificate...")
	issueTx, err := certificateRegistry.IssueCertificate(
		auth,
		certHash,
	)
	if err != nil {
		log.Fatal("Failed to issue certificate:", err)
	}

	receipt, err := bind.WaitMined(ctx, client, issueTx)
	if err != nil {
		log.Fatal("Failed waiting for certificate issue tx:", err)
	}

	fmt.Println("✅ Certificate issued successfully")
	fmt.Println("📦 Stored in block:", receipt.BlockNumber)
}
