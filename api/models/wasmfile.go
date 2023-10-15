package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMFile struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	DownloadLink string    `gorm:"type:text;not null"`
	Filename     string    `gorm:"type:text;not null"`
	Path         string    `gorm:"type:text;not null"`
	Sha256       string    `gorm:"type:varchar(255);not null"`
	Size         int       `gorm:"type:int;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ReleaseID    uuid.UUID `gorm:"type:uuid;not null"`
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (plugin *WASMFile) BeforeCreate(tx *gorm.DB) (err error) {
	plugin.ID = uuid.New() // Generate a new UUID for the record.
	return
}
