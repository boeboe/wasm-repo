package postgresrepo

import (
	"errors"
	"fmt"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	Database *gorm.DB
}

// CreateBinary saves the binary data to the PostgreSQL database.
func (r *PostgresRepository) CreateBinary(binary *sharedtypes.WASMBinary) error {
	if err := r.Database.Create(binary).Error; err != nil {
		return err
	}
	return nil
}

// GetBinaryByID retrieves the binary data by ID from the PostgreSQL database.
func (r *PostgresRepository) GetBinaryByID(id uint) (*sharedtypes.WASMBinary, error) {
	var binary sharedtypes.WASMBinary
	if err := r.Database.First(&binary, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("binary with ID %d not found", id)
		}
		return nil, err
	}
	return &binary, nil
}
