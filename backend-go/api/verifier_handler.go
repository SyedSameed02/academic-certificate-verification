package api

import (
	"encoding/json"
	"net/http"

	"backend-go/blockchain"
)

type VerifyRequest struct {
	CertificateHash string `json:"certificate_hash"`
}

func VerifyCertificate(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	json.NewDecoder(r.Body).Decode(&req)

	valid, revoked := blockchain.VerifyCertificate(req.CertificateHash)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid":   valid,
		"revoked": revoked,
	})
}
