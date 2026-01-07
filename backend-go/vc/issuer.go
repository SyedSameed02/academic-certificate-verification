package vc

import (
	"encoding/json"

	"backend-go/models"
	"backend-go/utils"
)

func IssueVC(
	issuerDID string,
	subjectDID string,
	cert models.Certificate,
) models.VerifiableCredential {

	raw, _ := json.Marshal(cert)
	hash := utils.HashStruct(string(raw))

	return models.VerifiableCredential{
		ID:         utils.HashStruct(subjectDID + hash),
		IssuerDID:  issuerDID,
		SubjectDID: subjectDID,
		Credential: cert,
		Hash:       hash,
	}
}
