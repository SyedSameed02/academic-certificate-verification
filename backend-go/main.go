package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	router := NewRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
