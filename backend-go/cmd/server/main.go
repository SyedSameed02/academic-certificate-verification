package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"academic-certificate-verification-backend/blockchain"
	"academic-certificate-verification-backend/internal/certificate"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractsConfig represents deployed smart contract addresses
// stored in config/contracts.json
type ContractsConfig struct {
	DIDRegistry         string `json:"didRegistry"`
	CertificateRegistry string `json:"certificateRegistry"`
}

func main() {
	// Context is used to control blockchain calls
	// It allows cancellation and timeouts if needed
	ctx := context.Background()

	// ------------------------------------------------------------------
	// STEP 1: Connect to local Hardhat blockchain
	// ------------------------------------------------------------------
	// The backend communicates with the blockchain using an RPC endpoint
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	fmt.Println("Connected to Ethereum node")

	// ------------------------------------------------------------------
	// STEP 2: Read deployed contract addresses
	// ------------------------------------------------------------------
	// Smart contract addresses are stored after deployment
	// This avoids hardcoding addresses in code
	configBytes, err := os.ReadFile("config/contracts.json")
	if err != nil {
		log.Fatal("Failed to read contracts.json:", err)
	}

	var config ContractsConfig
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatal("Failed to parse contracts.json:", err)
	}

	// ------------------------------------------------------------------
	// STEP 3: Create certificate data (off-chain)
	// ------------------------------------------------------------------
	// This represents the original academic certificate information
	// Blockchain never stores these fields directly
	cert := certificate.CertificateData{
		StudentName: "Rahul Sharma",
		RollNumber:  "CSE2021012",
		Course:      "B.Tech Computer Science",
		University:  "MITS",
		Year:        "2025",
	}

	// ------------------------------------------------------------------
	// STEP 4: Hash the certificate
	// ------------------------------------------------------------------
	// Hashing ensures:
	// - Integrity (any change alters the hash)
	// - Privacy (raw data is not stored on-chain)
	// The same data always produces the same hash
	certHash := certificate.HashCertificate(cert)

	// ------------------------------------------------------------------
	// STEP 5: Load CertificateRegistry smart contract
	// ------------------------------------------------------------------
	// This contract stores and verifies certificate hashes
	certificateRegistryAddress := common.HexToAddress(config.CertificateRegistry)

	certificateRegistry, err := blockchain.NewCertificateRegistry(
		certificateRegistryAddress,
		client,
	)
	if err != nil {
		log.Fatal("Failed to load CertificateRegistry contract:", err)
	}

	// ------------------------------------------------------------------
	// STEP 6: Call verification function (read-only)
	// ------------------------------------------------------------------
	// CallOpts is used for view/pure contract calls
	// No gas is required because this does not modify blockchain state
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	// The smart contract returns factual data only:
	// - exists  : whether the certificate hash was issued
	// - issuer  : address of issuing institution
	// - revoked : whether the certificate was invalidated
	exists, err := certificateRegistry.VerifyCertificate(
		callOpts,
		certHash,
	)
	if err != nil {
		log.Fatal("Certificate verification failed:", err)
	}

	// ------------------------------------------------------------------
	// STEP 7: Interpret verification result (business logic)
	// ------------------------------------------------------------------
	// Blockchain provides facts
	// Backend converts facts into a trust decision
	fmt.Println("Certificate Verification Result")

	if !exists {
		fmt.Println("Certificate is INVALID: not found on blockchain")
		return
	}

	

	fmt.Println("Certificate is VALID")
	//fmt.Println("Issued by institution address:", issuer.Hex())
}
