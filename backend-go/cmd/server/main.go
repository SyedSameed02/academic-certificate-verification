package main

import (
	"log"
	"net/http"

	"backend-go/api"
)

func main() {
	router := api.SetupRoutes()

	log.Println("🚀 Server running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
