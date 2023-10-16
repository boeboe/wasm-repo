package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/boeboe/wasm-repo/api/middleware"
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
		panic(err)
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
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wasmFile)
}

// DownloadFileHandler handles the request to download a WASMFile for internal purposes.
func (h *WASMFileHandler) DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileID, _ := r.Context().Value(middleware.FileIDKey).(uuid.UUID)
	wasmFile, err := h.Repo.GetFileByID(fileID)
	if err != nil {
		panic(err)
	}
	content, err := h.Repo.RetrieveFileContent(wasmFile.Filename)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+wasmFile.Filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}

// ConsumeFileHandler handles the request to download a WASMFile for external consumption purposes.
func (h *WASMFileHandler) ConsumeFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	downloadAlias, _ := vars["downloadAlias"]
	wasmFile, err := h.Repo.GetFileByDownloadAlias(downloadAlias)
	if err != nil {
		panic(err)
	}
	content, err := h.Repo.RetrieveFileContent(wasmFile.Filename)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+wasmFile.Filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}
