package repository

import (
	"github.com/boeboe/wasm-repo/api/errors"
	"gorm.io/gorm"
)

// BaseRepo contains the common methods and logic for all repositories.
type BaseRepo struct {
	Database *gorm.DB
}

func (r *BaseRepo) wrapDBError(operation string, err error) error {
	if err != nil {
		return &errors.DatabaseError{
			Operation: operation,
			Err:       err,
		}
	}
	return nil
}

func (r *BaseRepo) wrapFileError(file string, operation string, err error) error {
	if err != nil {
		return &errors.FileError{
			File:      file,
			Operation: operation,
			Err:       err,
		}
	}
	return nil
}
