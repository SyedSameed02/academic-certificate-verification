package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashStruct(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}
