package certificate

import (
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func HashCertificate(cert CertificateData) [32]byte {
	data := strings.Join([]string{
		cert.StudentName,
		cert.RollNumber,
		cert.Course,
		cert.University,
		cert.Year,
	}, "|")

	return crypto.Keccak256Hash([]byte(data))
}
