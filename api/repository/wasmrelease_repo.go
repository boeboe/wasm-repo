package repository

import (
	"github.com/boeboe/wasm-repo/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMReleaseRepo struct {
	Database *gorm.DB
}

// ListAllReleasesForPlugin lists all WASMRelease entries for a given WASMPlugin from the PostgreSQL database.
func (r *WASMReleaseRepo) ListAllReleasesForPlugin(pluginID uuid.UUID) ([]models.WASMRelease, error) {
	var releases []models.WASMRelease
	if err := r.Database.Where("plugin_id = ?", pluginID).Find(&releases).Error; err != nil {
		return nil, err
	}
	return releases, nil
}

// CreateReleaseForPlugin creates a new WASMRelease for a specific WASMPlugin in the PostgreSQL database.
func (r *WASMReleaseRepo) CreateReleaseForPlugin(pluginID uuid.UUID, release *models.WASMRelease) error {
	release.PluginID = pluginID
	return r.Database.Create(release).Error
}

// GetReleaseByID retrieves a specific WASMRelease by its ID for a given WASMPlugin from the PostgreSQL database.
func (r *WASMReleaseRepo) GetReleaseByID(pluginID uuid.UUID, releaseID uuid.UUID) (*models.WASMRelease, error) {
	var release models.WASMRelease
	if err := r.Database.Where("plugin_id = ? AND id = ?", pluginID, releaseID).First(&release).Error; err != nil {
		return nil, handleDBError(err)
	}
	return &release, nil
}

// UpdateReleaseForPlugin updates a specific WASMRelease for a given WASMPlugin in the PostgreSQL database.
func (r *WASMReleaseRepo) UpdateReleaseForPlugin(pluginID uuid.UUID, release *models.WASMRelease) error {
	release.PluginID = pluginID
	return r.Database.Save(release).Error
}

// DeleteReleaseForPlugin deletes a specific WASMRelease by its ID for a given WASMPlugin from the PostgreSQL database.
func (r *WASMReleaseRepo) DeleteReleaseForPlugin(pluginID uuid.UUID, releaseID uuid.UUID) error {
	return r.Database.Delete(&models.WASMRelease{}, "plugin_id = ? AND id = ?", pluginID, releaseID).Error
}
