package zkp

import "backend-go/utils"

func GenerateProof(hash string) string {
	return utils.HashStruct("zkp:" + hash)
}

func VerifyProof(hash, proof string) bool {
	return GenerateProof(hash) == proof
}
