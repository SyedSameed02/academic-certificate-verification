package handlers

import (
	"encoding/json"
	"net/http"
	"math/big"

	"backend-go/blockchain"

	"github.com/ethereum/go-ethereum/common"
)

type IssueRequest struct {
	DegreeHash          string `json:"degreeHash"`
	CGPA                int64  `json:"cgpa"`
	IssuerDidHash       string `json:"issuerDidHash"`
	IssuerSignatureHash string `json:"issuerSignatureHash"`
}

type IssueResponse struct {
	Status          string `json:"status"`
	CertificateHash string `json:"certificateHash"`
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	var req IssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// --- Parse numeric fields ---
	degreeHash, ok := new(big.Int).SetString(req.DegreeHash, 10)
	if !ok {
		http.Error(w, "invalid degreeHash", http.StatusBadRequest)
		return
	}

	issuerDidHash, ok := new(big.Int).SetString(req.IssuerDidHash, 10)
	if !ok {
		http.Error(w, "invalid issuerDidHash", http.StatusBadRequest)
		return
	}

	issuerSigHash, ok := new(big.Int).SetString(req.IssuerSignatureHash, 10)
	if !ok {
		http.Error(w, "invalid issuerSignatureHash", http.StatusBadRequest)
		return
	}

	cgpa := big.NewInt(req.CGPA)

	// --- Poseidon hash (same as ZKP) ---
	poseidonHash, err := blockchain.PoseidonHash(
		degreeHash,
		cgpa,
		issuerDidHash,
		issuerSigHash,
	)
	if err != nil {
		http.Error(w, "hashing failed", http.StatusInternalServerError)
		return
	}

	// --- Convert to bytes32 ---
	certHash := common.BytesToHash(
		poseidonHash.FillBytes(make([]byte, 32)),
	)

	// --- Store on blockchain ---
	client := blockchain.NewClient()
	certSvc := blockchain.NewCertificateService(client)

	if err := certSvc.IssueCertificate(certHash); err != nil {
		http.Error(w, "blockchain issue failed", http.StatusInternalServerError)
		return
	}

	// --- Success ---
	respond(w, IssueResponse{
		Status:          "issued",
		CertificateHash: certHash.Hex(),
	})
}
