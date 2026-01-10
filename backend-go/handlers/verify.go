package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"math/big"

	"backend-go/blockchain"
	"backend-go/zkp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type VerifyRequest struct {
	Proof  json.RawMessage `json:"proof"`
	Public []string        `json:"public"`
}

type VerifyResponse struct {
	Verified bool   `json:"verified"`
	Reason   string `json:"reason,omitempty"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// --- Create unique temp files ---
	id := uuid.New().String()

	proofPath := filepath.Join(os.TempDir(), "proof_"+id+".json")
	publicPath := filepath.Join(os.TempDir(), "public_"+id+".json")

	// Ensure cleanup
	defer os.Remove(proofPath)
	defer os.Remove(publicPath)

	// --- Write proof ---
	if err := os.WriteFile(proofPath, req.Proof, 0644); err != nil {
		http.Error(w, "failed to save proof", http.StatusInternalServerError)
		return
	}

	// --- Write public inputs ---
	publicBytes, err := json.Marshal(req.Public)
	if err != nil {
		http.Error(w, "invalid public inputs", http.StatusBadRequest)
		return
	}

	if err := os.WriteFile(publicPath, publicBytes, 0644); err != nil {
		http.Error(w, "failed to save public inputs", http.StatusInternalServerError)
		return
	}

	// --- ZKP verification ---
	valid, err := zkp.VerifyProof(proofPath, publicPath)
	if err != nil || !valid {
		respond(w, VerifyResponse{
			Verified: false,
			Reason:   "invalid_proof",
		})
		return
	}

	// --- Extract on-chain hash ---
	if len(req.Public) < 1 {
		respond(w, VerifyResponse{
			Verified: false,
			Reason:   "invalid_public_inputs",
		})
		return
	}

	poseidonInt, ok := new(big.Int).SetString(req.Public[0], 10)
	if !ok {
		respond(w, VerifyResponse{
			Verified: false,
			Reason:   "invalid_hash_format",
		})
		return
	}

	// Convert big.Int → bytes32
	certHash := common.BytesToHash(
		poseidonInt.FillBytes(make([]byte, 32)),
	)

	// --- Blockchain checks ---
	client := blockchain.NewClient()
	certSvc := blockchain.NewCertificateService(client)

	exists, err := certSvc.Exists(certHash)
	if err != nil || !exists {
		respond(w, VerifyResponse{
			Verified: false,
			Reason:   "certificate_not_found",
		})
		return
	}

	_, revoked, _, err := certSvc.GetCertificate(certHash)
	if err != nil || revoked {
		respond(w, VerifyResponse{
			Verified: false,
			Reason:   "certificate_revoked",
		})
		return
	}

	// --- VERIFIED ---
	respond(w, VerifyResponse{
		Verified: true,
	})
}

func respond(w http.ResponseWriter, resp any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
