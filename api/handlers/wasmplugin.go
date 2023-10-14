package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/boeboe/wasm-repo/api/models"
	"github.com/boeboe/wasm-repo/api/models/entities"
	"github.com/google/uuid"
)

type WASMPluginHandler struct {
	Repo *models.WASMRepository
}

// ListAllPluginsHandler handles the request to list all WASMPlugins
func (h *WASMPluginHandler) ListAllPluginsHandler(w http.ResponseWriter, r *http.Request) {
	plugins, err := h.Repo.ListAllPlugins()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(plugins)
}

// CreatePluginHandler handles the request to create a new WASMPlugin
func (h *WASMPluginHandler) CreatePluginHandler(w http.ResponseWriter, r *http.Request) {
	var plugin entities.WASMPlugin
	if err := json.NewDecoder(r.Body).Decode(&plugin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.CreatePlugin(&plugin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plugin)
}

// GetPluginByIDHandler handles the request to get a specific WASMPlugin by its ID
func (h *WASMPluginHandler) GetPluginByIDHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	plugin, err := h.Repo.GetPluginByID(pluginID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(plugin)
}

// UpdatePluginHandler handles the request to update a specific WASMPlugin
func (h *WASMPluginHandler) UpdatePluginHandler(w http.ResponseWriter, r *http.Request) {
	var plugin entities.WASMPlugin
	if err := json.NewDecoder(r.Body).Decode(&plugin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	plugin.ID = pluginID
	if err := h.Repo.UpdatePlugin(&plugin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plugin)
}

// DeletePluginHandler handles the request to delete a specific WASMPlugin
func (h *WASMPluginHandler) DeletePluginHandler(w http.ResponseWriter, r *http.Request) {
	pluginID, _ := r.Context().Value("pluginID").(uuid.UUID)
	if err := h.Repo.DeletePlugin(pluginID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
