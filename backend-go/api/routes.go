package api

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()

	// Issuer
	api.HandleFunc("/issuer/issue", IssueCertificate).Methods("POST")
	api.HandleFunc("/issuer/revoke", RevokeCertificate).Methods("POST")
	api.HandleFunc("/issuer/certificates/{did}", GetIssuedCertificates).Methods("GET")

	// Student
	api.HandleFunc("/student/credentials/{did}", GetStudentCredentials).Methods("GET")
	api.HandleFunc("/student/share", ShareCredential).Methods("POST")

	// Verifier
	api.HandleFunc("/verifier/verify", VerifyCertificate).Methods("POST")
	api.HandleFunc("/verifier/verify-zkp", VerifyWithZKP).Methods("POST")

	// DID
	api.HandleFunc("/did/create", CreateDID).Methods("POST")
	api.HandleFunc("/did/resolve/{did}", ResolveDID).Methods("GET")
}
