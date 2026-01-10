package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"backend-go/handlers"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/issue", handlers.IssueHandler).Methods("POST")
	api.HandleFunc("/revoke", handlers.RevokeHandler).Methods("POST")
	api.HandleFunc("/verify", handlers.VerifyHandler).Methods("POST")

	return r
}
