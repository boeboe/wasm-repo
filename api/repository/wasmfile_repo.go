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
	fmt.Printf("CreateFile: %+v \n\n", file)
	err := r.Database.Create(file).Error
	return r.wrapDBError("CreateFile", err)
}

// GetFileByID retrieves a WASMFile by its ID.
func (r *WASMFileRepo) GetFileByID(id uuid.UUID) (*models.WASMFile, error) {
	var file models.WASMFile
	err := r.Database.First(&file, "id = ?", id).Error
	return &file, r.wrapDBError("GetFileByID", err)
}

// UpdateFile updates a WASMFile record in the database.
func (r *WASMFileRepo) UpdateFile(file *models.WASMFile) error {
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
