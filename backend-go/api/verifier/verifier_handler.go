package verifier

import (
	"encoding/json"
	"net/http"

	"backend-go/zkp"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	proof := r.URL.Query().Get("proof")

	valid := zkp.VerifyProof(hash, proof)

	json.NewEncoder(w).Encode(map[string]bool{
		"valid": valid,
	})
}
