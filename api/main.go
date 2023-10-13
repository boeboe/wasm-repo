package main

import (
	"log"
	"net/http"

	"github.com/boeboe/wasm-repo/api/handlers"
	"github.com/boeboe/wasm-repo/api/models"
)

func main() {
	models.ConnectDatabase()

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to this wasm-repo"))
	})

	// Upload handler
	http.HandleFunc("/upload", handlers.CreateWASMBinaryHandler)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
