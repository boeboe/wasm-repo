package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/boeboe/wasm-repo/api/models"
	"github.com/boeboe/wasm-repo/api/models/entities"
	"github.com/gorilla/mux"
)

type WASMReleaseHandler struct {
	Repo *models.WASMRepository
}

// ListAllReleasesForPluginHandler handles the request to list all WASMReleases for a specific WASMPlugin
func (h *WASMReleaseHandler) ListAllReleasesForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID := mux.Vars(r)["pluginID"]
	releases, err := h.Repo.ListAllReleasesForPlugin(pluginID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(releases)
}

// CreateReleaseForPluginHandler handles the request to create a new WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) CreateReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID := mux.Vars(r)["pluginID"]
	var release entities.WASMRelease
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.CreateReleaseForPlugin(pluginID, &release); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(release)
}

// GetReleaseByIDHandler handles the request to get a specific WASMRelease by its ID for a specific WASMPlugin
func (h *WASMReleaseHandler) GetReleaseByIDHandler(w http.ResponseWriter, r *http.Request) {
	pluginID := mux.Vars(r)["pluginID"]
	releaseID := mux.Vars(r)["releaseID"]
	release, err := h.Repo.GetReleaseByID(pluginID, releaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(release)
}

// UpdateReleaseForPluginHandler handles the request to update a specific WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) UpdateReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID := mux.Vars(r)["pluginID"]
	var release entities.WASMRelease
	if err := json.NewDecoder(r.Body).Decode(&release); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.UpdateReleaseForPlugin(pluginID, &release); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(release)
}

// DeleteReleaseForPluginHandler handles the request to delete a specific WASMRelease for a specific WASMPlugin
func (h *WASMReleaseHandler) DeleteReleaseForPluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID := mux.Vars(r)["pluginID"]
	releaseID := mux.Vars(r)["releaseID"]
	if err := h.Repo.DeleteReleaseForPlugin(pluginID, releaseID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
