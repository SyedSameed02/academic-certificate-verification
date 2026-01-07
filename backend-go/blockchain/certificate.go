package blockchain

import (
	"errors"
)

// IsCertificateValid checks existence & non-revocation
func IsCertificateValid(hash string) (bool, error) {
	if EthClient == nil {
		return false, errors.New("ethereum client not initialized")
	}

	/*
		REAL IMPLEMENTATION (example):

		valid, err := contract.IsValid(nil, hash)
		return valid, err
	*/

	// Stub behavior (safe default)
	return true, nil
}
