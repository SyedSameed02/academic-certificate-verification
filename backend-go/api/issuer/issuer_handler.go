package issuer

import (
	"encoding/json"
	"net/http"

	"backend-go/did"
	"backend-go/ipfs"
	"backend-go/models"
	"backend-go/vc"
	"backend-go/blockchain"
)

func IssueCertificate(w http.ResponseWriter, r *http.Request) {
	var cert models.Certificate
	json.NewDecoder(r.Body).Decode(&cert)

	issuerDID := did.GenerateDID()
	studentDID := did.GenerateDID()

	vcObj := vc.IssueVC(issuerDID, studentDID, cert)

	data, _ := json.Marshal(vcObj)
	ipfsHash, _ := ipfs.Store(data)
	vcObj.IPFSHash = ipfsHash

	blockchain.IssueCertificate(vcObj.Hash)

	json.NewEncoder(w).Encode(vcObj)
}

/*
POST /api/issuer/revoke
Body:
{
	"hash": "0xabc123hash"
}

Revokes a certificate by hash.
Only issuer performs on-chain revocation.
*/


//REMOVED REVOCATION LOGIC SO COMMENTED 

// func RevokeCertificate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var req struct {
// 		Hash string `json:"hash"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	if req.Hash == "" {
// 		http.Error(w, "certificate hash is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Call blockchain layer (abigen-backed)
// 	if err := blockchain.RevokeCertificate(req.Hash); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"status": "revoked",
// 		"hash":   req.Hash,
// 	})
// }
