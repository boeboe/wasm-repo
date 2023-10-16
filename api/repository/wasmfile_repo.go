package repository

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boeboe/wasm-repo/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMFileRepo struct {
	Database *gorm.DB
	BaseRepo
}

// CreateFile creates a new WASMFile record in the database.
func (r *WASMFileRepo) CreateFile(file *models.WASMFile) error {
	if err := r.setDownloadAliasForFile(file); err != nil {
		return err
	}
	err := r.Database.Create(file).Error
	return r.wrapDBError("CreateFile", err)
}

// GetFileByID retrieves a WASMFile by its ID.
func (r *WASMFileRepo) GetFileByID(id uuid.UUID) (*models.WASMFile, error) {
	var file models.WASMFile
	err := r.Database.First(&file, "id = ?", id).Error
	return &file, r.wrapDBError("GetFileByID", err)
}

// GetFileByReleaseID retrieves a WASMFile associated with a specific WASMRelease.
func (r *WASMFileRepo) GetFileByReleaseID(releaseID uuid.UUID) (*models.WASMFile, error) {
	var file models.WASMFile
	err := r.Database.First(&file, "release_id = ?", releaseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.wrapDBError("GetFileByReleaseID", fmt.Errorf("no file found for release ID %s", releaseID))
		}
		return nil, r.wrapDBError("GetFileByReleaseID", err)
	}
	return &file, nil
}

// GetFileByDownloadAlias retrieves a WASMFile by its DownloadAlias.
func (r *WASMFileRepo) GetFileByDownloadAlias(downloadAlias string) (*models.WASMFile, error) {
	var file models.WASMFile
	err := r.Database.First(&file, "download_alias = ?", downloadAlias).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.wrapDBError("GetFileByDownloadAlias", fmt.Errorf("no file found for download alias %s", downloadAlias))
		}
		return nil, r.wrapDBError("GetFileByDownloadAlias", err)
	}
	return &file, nil
}

// UpdateFile updates a WASMFile record in the database.
func (r *WASMFileRepo) UpdateFile(file *models.WASMFile) error {
	if err := r.setDownloadAliasForFile(file); err != nil {
		return err
	}
	err := r.Database.Save(file).Error
	return r.wrapDBError("UpdateFile", err)
}

// DeleteFile deletes a WASMFile record from the database.
func (r *WASMFileRepo) DeleteFile(id uuid.UUID) error {
	err := r.Database.Delete(&models.WASMFile{}, "id = ?", id).Error
	return r.wrapDBError("DeleteFile", err)
}

// StoreFileContent saves the file content to a storage location and returns the path.
// For simplicity, we're storing it on the local filesystem, but this can be extended to cloud storage.
func (r *WASMFileRepo) StoreFileContent(filename string, content []byte) (string, error) {
	storagePath := "./output/" // Define your storage directory
	fullPath := filepath.Join(storagePath, filename)

	err := os.WriteFile(fullPath, content, 0644)
	return fullPath, r.wrapFileError(filename, "StoreFileContent", err)
}

// RetrieveFileContent retrieves the file content from a storage location given a filename.
func (r *WASMFileRepo) RetrieveFileContent(filename string) ([]byte, error) {
	storagePath := "./output/" // Define your storage directory
	fullPath := filepath.Join(storagePath, filename)

	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		return nil, r.wrapFileError(filename, "RetrieveFileContent", fmt.Errorf("file %s does not exist", filename))
	}

	content, err := os.ReadFile(fullPath)
	return content, r.wrapFileError(filename, "RetrieveFileContent", err)
}

// DeleteFileContent deletes the file from a storage location given a filename.
func (r *WASMFileRepo) DeleteFileContent(filename string) error {
	storagePath := "./output/" // Define your storage directory
	fullPath := filepath.Join(storagePath, filename)

	if err := os.Remove(fullPath); err != nil {
		return r.wrapFileError(filename, "DeleteFileContent", err)
	}
	return nil
}

// setDownloadLinkForFile sets the download link for a given WASMFile based on its associated plugin and release.
func (r *WASMFileRepo) setDownloadAliasForFile(file *models.WASMFile) error {
	var release models.WASMRelease
	err := r.Database.First(&release, file.ReleaseID).Error
	if err != nil {
		return err
	}
	var plugin models.WASMPlugin
	err = r.Database.First(&plugin, release.PluginID).Error
	if err != nil {
		return err
	}
	file.SetDownloadAlias(plugin.Name, release.Version)
	return nil
}
