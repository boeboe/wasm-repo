package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMRelease struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Version     string    `gorm:"type:varchar(255);not null"`
	Sha256      string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Size        int       `gorm:"type:int;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Location    WASMLocation `gorm:"foreignKey:ReleaseID"`
	PluginID    uuid.UUID    `gorm:"type:uuid;not null"`
}

type WASMLocation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type      string    `gorm:"type:varchar(255);not null"`
	Location  string    `gorm:"type:text;not null"`
	ReleaseID uuid.UUID `gorm:"type:uuid;not null"`
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (plugin *WASMRelease) BeforeCreate(tx *gorm.DB) (err error) {
	plugin.ID = uuid.New() // Generate a new UUID for the record.
	return
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (plugin *WASMLocation) BeforeCreate(tx *gorm.DB) (err error) {
	plugin.ID = uuid.New() // Generate a new UUID for the record.
	return
}
