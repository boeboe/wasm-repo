package main

import (
	"log"
	"net/http"

	"github.com/boeboe/wasm-repo/api/models"
)

func main() {
	models.ConnectDatabase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WASM Repo API"))
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
