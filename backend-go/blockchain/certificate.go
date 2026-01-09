package blockchain

import (
	"errors"
	"fmt"
)

func IsCertificateValid(hash string) (bool, error) {
	if EthClient == nil {
		return false, errors.New("ethereum client not initialized")
	}

	
	valid, err := VerifyCertificate(hash)
	if err != nil {
		return false, err
	}
	fmt.Sprint(valid)

	return true, nil
}
