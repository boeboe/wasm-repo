package models

import (
	"errors"
	"fmt"

	"github.com/boeboe/wasm-repo/api/models/entities"
	"gorm.io/gorm"
)

type WASMRepository struct {
	Database *gorm.DB
}

// ListAllPlugins lists all WASMPlugin entries from the PostgreSQL database.
func (r *WASMRepository) ListAllPlugins() ([]entities.WASMPlugin, error) {
	var plugins []entities.WASMPlugin
	if err := r.Database.Find(&plugins).Error; err != nil {
		return nil, err
	}
	return plugins, nil
}

// CreatePlugin creates a new WASMPlugin entry in the PostgreSQL database.
func (r *WASMRepository) CreatePlugin(plugin *entities.WASMPlugin) error {
	return r.Database.Create(plugin).Error
}

// GetPluginByID retrieves a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMRepository) GetPluginByID(id string) (*entities.WASMPlugin, error) {
	var plugin entities.WASMPlugin
	if err := r.Database.First(&plugin, "id = ?", id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return &plugin, nil
}

// UpdatePlugin updates a specific WASMPlugin entry in the PostgreSQL database.
func (r *WASMRepository) UpdatePlugin(plugin *entities.WASMPlugin) error {
	return r.Database.Save(plugin).Error
}

// DeletePlugin deletes a specific WASMPlugin by its ID from the PostgreSQL database.
func (r *WASMRepository) DeletePlugin(id string) error {
	return r.Database.Delete(&entities.WASMPlugin{}, "id = ?", id).Error
}

// ListAllReleasesForPlugin lists all WASMRelease entries for a given WASMPlugin from the PostgreSQL database.
func (r *WASMRepository) ListAllReleasesForPlugin(pluginID string) ([]entities.WASMRelease, error) {
	var releases []entities.WASMRelease
	if err := r.Database.Where("plugin_id = ?", pluginID).Find(&releases).Error; err != nil {
		return nil, err
	}
	return releases, nil
}

// CreateReleaseForPlugin creates a new WASMRelease for a specific WASMPlugin in the PostgreSQL database.
func (r *WASMRepository) CreateReleaseForPlugin(pluginID string, release *entities.WASMRelease) error {
	release.PluginID = pluginID
	return r.Database.Create(release).Error
}

// GetReleaseByID retrieves a specific WASMRelease by its ID for a given WASMPlugin from the PostgreSQL database.
func (r *WASMRepository) GetReleaseByID(pluginID string, releaseID string) (*entities.WASMRelease, error) {
	var release entities.WASMRelease
	if err := r.Database.Where("plugin_id = ? AND id = ?", pluginID, releaseID).First(&release).Error; err != nil {
		return nil, handleDBError(err)
	}
	return &release, nil
}

// UpdateReleaseForPlugin updates a specific WASMRelease for a given WASMPlugin in the PostgreSQL database.
func (r *WASMRepository) UpdateReleaseForPlugin(pluginID string, release *entities.WASMRelease) error {
	release.PluginID = pluginID
	return r.Database.Save(release).Error
}

// DeleteReleaseForPlugin deletes a specific WASMRelease by its ID for a given WASMPlugin from the PostgreSQL database.
func (r *WASMRepository) DeleteReleaseForPlugin(pluginID string, releaseID string) error {
	return r.Database.Delete(&entities.WASMRelease{}, "plugin_id = ? AND id = ?", pluginID, releaseID).Error
}

// handleDBError is a helper function to handle common database errors.
func handleDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("record not found: %v", err)
	}
	return err
}
