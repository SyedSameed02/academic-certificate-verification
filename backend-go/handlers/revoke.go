package handlers

import (
	"encoding/json"
	"net/http"

	"backend-go/blockchain"

	"github.com/ethereum/go-ethereum/common"
)

type RevokeRequest struct {
	CertificateHash string `json:"certificateHash"`
}

type RevokeResponse struct {
	Status string `json:"status"`
}

func RevokeHandler(w http.ResponseWriter, r *http.Request) {
	var req RevokeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}






	certHash := common.HexToHash(req.CertificateHash)
	if certHash == (common.Hash{}) {
		http.Error(w, "invalid certificate hash", http.StatusBadRequest)
		return
	}


	// --- Blockchain revoke ---
	client := blockchain.NewClient()
	certSvc := blockchain.NewCertificateService(client)



	if err := certSvc.RevokeCertificate(certHash); err != nil {
		http.Error(w, "blockchain revoke failed", http.StatusInternalServerError)
		return
	}

	respond(w, RevokeResponse{
		Status: "revoked",
	})
}
