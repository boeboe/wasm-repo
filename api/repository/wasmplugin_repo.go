package repository

import (
	"github.com/boeboe/wasm-repo/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMPluginRepo struct {
	Database *gorm.DB
}

// ListAllPlugins lists all WASMPlugin entries from the PostgreSQL database.
func (r *WASMPluginRepo) ListAllPlugins() ([]models.WASMPlugin, error) {
	var plugins []models.WASMPlugin
	if err := r.Database.Find(&plugins).Error; err != nil {
		return nil, err
	}
	return plugins, nil
}

// CreatePlugin creates a new WASMPlugin entry in the PostgreSQL database.
func (r *WASMPluginRepo) CreatePlugin(plugin *models.WASMPlugin) error {
	return r.Database.Create(plugin).Error
}

// GetPluginByID retrieves a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMPluginRepo) GetPluginByID(id uuid.UUID) (*models.WASMPlugin, error) {
	var plugin models.WASMPlugin
	if err := r.Database.First(&plugin, "id = ?", id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return &plugin, nil
}

// UpdatePlugin updates a specific WASMPlugin entry in the PostgreSQL database.
func (r *WASMPluginRepo) UpdatePlugin(plugin *models.WASMPlugin) error {
	return r.Database.Save(plugin).Error
}

// DeletePlugin deletes a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMPluginRepo) DeletePlugin(id uuid.UUID) error {
	return r.Database.Delete(&models.WASMPlugin{}, "id = ?", id).Error
}
