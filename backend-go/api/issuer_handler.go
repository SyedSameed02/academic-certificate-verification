package api

import (
	"encoding/json"
	"net/http"

	"/blockchain"
	"internal"
)

func IssueCertificate(w http.ResponseWriter, r *http.Request) {
	var cert internal.Certificate

	if err := json.NewDecoder(r.Body).Decode(&cert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	txHash, err := blockchain.IssueCertificate(cert)
	if err != nil {
		http.Error(w, "Blockchain issue failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "issued",
		"txHash": txHash,
	})
}

