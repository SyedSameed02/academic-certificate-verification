package blockchain

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

func PoseidonHash(inputs ...*big.Int) (*big.Int, error) {
	return poseidon.Hash(inputs)
}
