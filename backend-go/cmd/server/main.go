package main

import (
	"log"

	"backend-go/blockchain"
	"backend-go/api"
	"backend-go/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	err = blockchain.Init(
		cfg.BlockChain.RPCUrl,
		cfg.BlockChain.CertificateRegistry,
		cfg.BlockChain.DIDRegistry,
	)
	if err != nil {
		log.Fatal("❌ Blockchain initialization failed:", err)
	}

	log.Println("🚀 Backend starting on port", cfg.Server.Port)

	err = api.StartServer(cfg)
	if err != nil {
		log.Fatal("❌ Server stopped:", err)
	}
}
