package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"backend-go/blockchain"
    "backend-go/internal/certificate"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractsConfig holds deployed contract addresses
// Loaded from config/contracts.json
type ContractsConfig struct {
	DIDRegistry         string `json:"didRegistry"`
	CertificateRegistry string `json:"certificateRegistry"`
}

func main() {
	// Context controls lifecycle of blockchain calls
	ctx := context.Background()

	// ------------------------------------------------------------------
	// STEP 1: Connect to local Hardhat blockchain
	// ------------------------------------------------------------------
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}

	fmt.Println("Connected to Ethereum node")

	// ------------------------------------------------------------------
	// STEP 2: Load deployed contract addresses
	// ------------------------------------------------------------------
	configBytes, err := os.ReadFile("config/contracts.json")
	if err != nil {
		log.Fatal("Failed to read contracts.json:", err)
	}

	var config ContractsConfig
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatal("Failed to parse contracts.json:", err)
	}

	// ------------------------------------------------------------------
	// STEP 3: Construct certificate data (off-chain)
	// ------------------------------------------------------------------
	// Certificate data is never stored on-chain
	cert := certificate.CertificateData{
		StudentName: "Rahul Sharma",
		RollNumber:  "CSE2021012",
		Course:      "B.Tech Computer Science",
		University:  "MITS",
		Year:        "2025",
	}

	// ------------------------------------------------------------------
	// STEP 4: Hash certificate
	// ------------------------------------------------------------------
	// Hash ensures integrity and privacy
	certHash := certificate.HashCertificate(cert)

	// ------------------------------------------------------------------
	// STEP 5: Load CertificateRegistry contract
	// ------------------------------------------------------------------
	certificateRegistryAddr := common.HexToAddress(config.CertificateRegistry)

	certificateRegistry, err := blockchain.NewCertificateRegistry(
		certificateRegistryAddr,
		client,
	)
	if err != nil {
		log.Fatal("Failed to load CertificateRegistry:", err)
	}

	// ------------------------------------------------------------------
	// STEP 6: Load DIDRegistry contract
	// ------------------------------------------------------------------
	didRegistryAddr := common.HexToAddress(config.DIDRegistry)

	didRegistry, err := blockchain.NewDIDRegistry(
		didRegistryAddr,
		client,
	)
	if err != nil {
		log.Fatal("Failed to load DIDRegistry:", err)
	}

	// ------------------------------------------------------------------
	// STEP 7: Perform read-only verification call
	// ------------------------------------------------------------------
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	// IMPORTANT:
	// VerifyCertificate returns (struct, error), NOT multiple values
	certInfo, err := certificateRegistry.VerifyCertificate(callOpts, certHash)
	if err != nil {
		log.Fatal("Blockchain verification failed:", err)
	}

	// ------------------------------------------------------------------
	// STEP 8: Interpret verification result
	// ------------------------------------------------------------------
	fmt.Println("Certificate Verification Result")

	// Case 1: Certificate never issued
	if !certInfo.Exists {
		fmt.Println("Certificate is INVALID: not found on blockchain")
		return
	}

	// Case 2: Certificate revoked
	if !certInfo.IsValid {
		fmt.Println("Certificate is INVALID: certificate has been revoked")
		return
	}

	// Case 3: Validate issuer using DIDRegistry
	isIssuerValid, err := didRegistry.IsValidIssuer(callOpts, certInfo.Issuer)
	if err != nil {
		log.Fatal("Issuer validation failed:", err)
	}

	if !isIssuerValid {
		fmt.Println("Certificate is INVALID: issuer is not registered")
		return
	}

	// Case 4: All checks passed
	fmt.Println("Certificate is VALID")
	fmt.Println("Issued by institution address:", certInfo.Issuer.Hex())
}
