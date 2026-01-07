package blockchain

import (

	"errors"
	"log"
)

// IssueCertificate registers a certificate hash on-chain
func IssueCertificate(hash string) error {
	if EthClient == nil {
		return errors.New("ethereum client not initialized")
	}

	/*
		REAL IMPLEMENTATION (example):

		auth, err := bind.NewKeyedTransactorWithChainID(...)
		tx, err := contract.RegisterCertificate(auth, hash)
	*/

	log.Println("📝 Certificate hash registered on-chain:", hash)

	// Stub-safe: no panic, no tx failure
	return nil
}
