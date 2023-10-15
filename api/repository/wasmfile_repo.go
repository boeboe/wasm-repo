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
}

// CreateFile creates a new WASMFile record in the database.
func (r *WASMFileRepo) CreateFile(file *models.WASMFile) error {
	fmt.Printf("CreateFile: %+v \n\n", file)
	return r.Database.Create(file).Error
}

// GetFileByID retrieves a WASMFile by its ID.
func (r *WASMFileRepo) GetFileByID(id uuid.UUID) (*models.WASMFile, error) {
	var file models.WASMFile
	if err := r.Database.First(&file, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

// UpdateFile updates a WASMFile record in the database.
func (r *WASMFileRepo) UpdateFile(file *models.WASMFile) error {
	return r.Database.Save(file).Error
}

// DeleteFile deletes a WASMFile record from the database.
func (r *WASMFileRepo) DeleteFile(id uuid.UUID) error {
	return r.Database.Delete(&models.WASMFile{}, "id = ?", id).Error
}

// StoreFileContent saves the file content to a storage location and returns the path.
// For simplicity, we're storing it on the local filesystem, but this can be extended to cloud storage.
func (r *WASMFileRepo) StoreFileContent(filename string, content []byte) (string, error) {
	storagePath := "./output/" // Define your storage directory
	fullPath := filepath.Join(storagePath, filename)

	if err := os.WriteFile(fullPath, content, 0644); err != nil {
		return "", err
	}
	return fullPath, nil
}

// RetrieveFileContent retrieves the file content from a storage location given a filename.
func (r *WASMFileRepo) RetrieveFileContent(filename string) ([]byte, error) {
	storagePath := "./output/" // Define your storage directory
	fullPath := filepath.Join(storagePath, filename)

	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
