package models

type VerifiableCredential struct {
	ID           string      `json:"id"`
	IssuerDID    string      `json:"issuer_did"`
	SubjectDID   string      `json:"subject_did"`
	Credential   Certificate `json:"credential"`
	Hash         string      `json:"hash"`
	IPFSHash     string      `json:"ipfs_hash"`
}
