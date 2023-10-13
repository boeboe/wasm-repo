// file api/handlers/wasm.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/boeboe/wasm-repo/api/models"
	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
)

func CreateWASMBinaryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Extract binary data
	file, _, err := r.FormFile("binary")
	if err != nil {
		http.Error(w, "Unable to get binary data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	binaryData := make([]byte, r.ContentLength)
	_, err = file.Read(binaryData)
	if err != nil {
		http.Error(w, "Unable to read binary data", http.StatusInternalServerError)
		return
	}

	// Extract metadata
	name := r.FormValue("name")
	description := r.FormValue("description")
	version := r.FormValue("version")

	binary := &sharedtypes.WASMBinary{
		Name:   name,
		Binary: binaryData,
		Metadata: sharedtypes.WASMMetadata{
			Description: description,
			Version:     version,
		},
	}

	err = models.Repo.CreateBinary(binary)
	if err != nil {
		http.Error(w, "Error saving binary", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Binary successfully uploaded")
}
