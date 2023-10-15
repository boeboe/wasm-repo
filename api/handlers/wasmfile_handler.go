package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/boeboe/wasm-repo/api/models"
	"github.com/boeboe/wasm-repo/api/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type WASMFileHandler struct {
	Repo *repository.WASMFileRepo
}

// UploadFileHandler handles the request to upload a WASMFile.
func (h *WASMFileHandler) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		http.Error(w, "Unable to parse file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Parse mandatory form fields "pluginID" and "releaseID"
	releaseID := r.FormValue("releaseID")

	if releaseID == "" {
		fmt.Printf("Error: Mandatory field 'releaseID' not provided\n")
		http.Error(w, "Mandatory field 'releaseID' not provided", http.StatusBadRequest)
		return
	}

	// Store the file content
	path, err := h.Repo.StoreFileContent(header.Filename, content)
	if err != nil {
		fmt.Printf("Error storing file: %v\n", err)
		http.Error(w, "Unable to store file", http.StatusInternalServerError)
		return
	}

	hash := sha256.Sum256(content)

	// Create a new WASMFile record
	wasmFile := &models.WASMFile{
		Filename:  header.Filename,
		Path:      path,
		Sha256:    hex.EncodeToString(hash[:]),
		Size:      len(content),
		ReleaseID: uuid.MustParse(releaseID),
	}
	if err := h.Repo.CreateFile(wasmFile); err != nil {
		fmt.Printf("Error creating file record: %v\n", err)
		http.Error(w, "Unable to create file record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wasmFile)
}

// DownloadFileHandler handles the request to download a WASMFile.
func (h *WASMFileHandler) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID, err := uuid.Parse(vars["fileID"])
	if err != nil {
		fmt.Printf("Invalid file ID error: %v\n", err)
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	// Retrieve the WASMFile record
	wasmFile, err := h.Repo.GetFileByID(fileID)
	if err != nil {
		fmt.Printf("File not found error: %v\n", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Retrieve the file content
	content, err := h.Repo.RetrieveFileContent(wasmFile.Filename)
	if err != nil {
		fmt.Printf("Error retrieving file: %v\n", err)
		http.Error(w, "Unable to retrieve file", http.StatusInternalServerError)
		return
	}

	// Set headers and write the file content to the response
	w.Header().Set("Content-Disposition", "attachment; filename="+wasmFile.Filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}
