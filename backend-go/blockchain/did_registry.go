package blockchain

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterIssuer(
	auth *bind.TransactOpts,
	contract *DIDRegistry,
	issuer common.Address,
	did string,
) {
	tx, err := contract.RegisterIssuer(auth, issuer, did)
	if err != nil {
		log.Fatal("Failed to register issuer:", err)
	}
	log.Println("Issuer registered, tx:", tx.Hash().Hex())
}

func IsValidIssuer(
	ctx context.Context,
	contract *DIDRegistry,
	issuer common.Address,
) bool {
	valid, err := contract.IsValidIssuer(&bind.CallOpts{Context: ctx}, issuer)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}
