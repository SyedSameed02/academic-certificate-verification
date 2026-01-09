package api

import (
	"net/http"

	"backend-go/api/issuer"
	"backend-go/api/student"
	"backend-go/api/verifier"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Issuer routes
	mux.HandleFunc("/api/issuer/issue", issuer.IssueCertificate)

	// Student routes
	mux.HandleFunc("/api/student/credentials/", student.GetCredentials)
	mux.HandleFunc("/api/student/share", student.ShareCredential)

	// Verifier routes
	mux.HandleFunc("/api/verifier/verify", verifier.Verify)

	return mux
}
