package blockchain

import (
	"backend-go/internal/certificate"
)

func RevokeCertificate(cert certificate.CertificateData) (string, error) {
	// call binding-level function
	return RevokeCertificate(cert)
}
