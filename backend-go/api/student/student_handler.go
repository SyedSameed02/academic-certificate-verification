package student

import (
	"encoding/json"
	"net/http"

	"backend-go/models"
	"backend-go/zkp"
)

/*
GET /api/student/credentials?did=did:edu:123
Returns credentials held by the student (off-chain reference)
*/
func GetCredentials(w http.ResponseWriter, r *http.Request) {
	did := r.URL.Query().Get("did")
	if did == "" {
		http.Error(w, "missing did", http.StatusBadRequest)
		return
	}

	// NOTE:
	// In a real system, this would come from:
	// - IPFS
	// - local wallet
	// - encrypted DB
	//
	// Here we return a stub to avoid breaking flow.

	response := []models.VerifiableCredential{
		{
			SubjectDID: did,
			Hash:       "0xabc123hash",
			IPFSHash:   "QmDummyIPFSHash",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
POST /api/student/share
Body:
{
	"hash": "0xabc123hash"
}

Returns ZKP-style proof without revealing certificate data
*/
func ShareCredential(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Hash string `json:"hash"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Hash == "" {
		http.Error(w, "missing hash", http.StatusBadRequest)
		return
	}

	proof := zkp.GenerateProof(req.Hash)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Proof{
		Hash:  req.Hash,
		Proof: proof,
	})
}
