package repository

import (
	"github.com/boeboe/wasm-repo/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMPluginRepo struct {
	Database *gorm.DB
	BaseRepo
}

// ListAllPlugins lists all WASMPlugin entries from the PostgreSQL database.
func (r *WASMPluginRepo) ListAllPlugins() ([]models.WASMPlugin, error) {
	var plugins []models.WASMPlugin
	err := r.Database.Find(&plugins).Error
	return plugins, r.wrapDBError("ListAllPlugins", err)
}

// SearchPluginsByName lists all WASMPlugin entries from the PostgreSQL database with an optional filter.
func (r *WASMPluginRepo) SearchPlugins(filter models.WASMPluginFilter) ([]models.WASMPlugin, error) {
	query := r.Database

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.Owner != "" {
		query = query.Where("owner = ?", filter.Owner)
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	var plugins []models.WASMPlugin
	err := query.Find(&plugins).Error
	return plugins, err
}

// CreatePlugin creates a new WASMPlugin entry in the PostgreSQL database.
func (r *WASMPluginRepo) CreatePlugin(plugin *models.WASMPlugin) error {
	err := r.Database.Create(plugin).Error
	return r.wrapDBError("CreatePlugin", err)
}

// GetPluginByID retrieves a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMPluginRepo) GetPluginByID(id uuid.UUID) (*models.WASMPlugin, error) {
	var plugin models.WASMPlugin
	err := r.Database.First(&plugin, "id = ?", id).Error
	return &plugin, r.wrapDBError("GetPluginByID", err)
}

// UpdatePlugin updates a specific WASMPlugin entry in the PostgreSQL database.
func (r *WASMPluginRepo) UpdatePlugin(plugin *models.WASMPlugin) error {
	err := r.Database.Save(plugin).Error
	return r.wrapDBError("UpdatePlugin", err)
}

// DeletePlugin deletes a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMPluginRepo) DeletePlugin(id uuid.UUID) error {
	err := r.Database.Delete(&models.WASMPlugin{}, "id = ?", id).Error
	return r.wrapDBError("DeletePlugin", err)
}
