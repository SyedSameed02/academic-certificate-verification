package did

import (
	"github.com/google/uuid"
)

func GenerateDID() string {
	return "did:edu:" + uuid.New().String()
}
