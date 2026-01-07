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

	blockchain.RegisterCertificate(vcObj.Hash)

	json.NewEncoder(w).Encode(vcObj)
}
