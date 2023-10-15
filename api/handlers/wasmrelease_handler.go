package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/boeboe/wasm-repo/api/errors"
	"github.com/boeboe/wasm-repo/api/models"
	"github.com/boeboe/wasm-repo/api/repository"
	"github.com/google/uuid"
)

type WASMReleaseHandler struct {
	FileRepo    *repository.WASMFileRepo
	ReleaseRepo *repository.WASMReleaseRepo
}

// ListAllReleasesForPluginHandler handles the request to list all WASMReleases for a specific WASMPlugin
func (h *WASMReleaseHandler) ListAllReleasesForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	releases, err := h.ReleaseRepo.ListAllReleasesForPlugin(pluginID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(releases)
}

// CreateReleaseForPluginHandler handles the request to create a new WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) CreateReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	var release models.WASMRelease
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		panic(&errors.JSONDecodingError{Source: "CreateReleaseForPluginHandler", Err: err})
	}
	if err := h.ReleaseRepo.CreateReleaseForPlugin(pluginID, &release); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(release)
}

// GetReleaseByIDHandler handles the request to get a specific WASMRelease by its ID for a specific WASMPlugin
func (h *WASMReleaseHandler) GetReleaseByIDHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	releaseID, _ := r.Context().Value("releaseID").(uuid.UUID)
	release, err := h.ReleaseRepo.GetReleaseByID(pluginID, releaseID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(release)
}

// UpdateReleaseForPluginHandler handles the request to update a specific WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) UpdateReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	releaseID, _ := r.Context().Value("releaseID").(uuid.UUID)
	var release models.WASMRelease
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		panic(&errors.JSONDecodingError{Source: "UpdateReleaseForPluginHandler", Err: err})
	}
	release.ID = releaseID
	if err := h.ReleaseRepo.UpdateReleaseForPlugin(pluginID, &release); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(release)
}

// DeleteReleaseForPluginHandler handles the request to delete a specific WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) DeleteReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	releaseID, _ := r.Context().Value("releaseID").(uuid.UUID)
	file, err := h.FileRepo.GetFileByReleaseID(releaseID)
	if err != nil {
		panic(err)
	}
	if err := h.FileRepo.DeleteFileContent(file.Filename); err != nil {
		panic(err)
	}
	if err := h.FileRepo.DeleteFile(file.ID); err != nil {
		panic(err)
	}
	if err := h.ReleaseRepo.DeleteReleaseForPlugin(pluginID, releaseID); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}
