package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMRelease struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Description string    `gorm:"type:text"`
	Version     string    `gorm:"type:varchar(255);not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	File        WASMFile  `gorm:"foreignKey:ReleaseID"`
	PluginID    uuid.UUID `gorm:"type:uuid;not null"`
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (wr *WASMRelease) BeforeCreate(tx *gorm.DB) (err error) {
	wr.ID = uuid.New() // Generate a new UUID for the record.
	return
}
