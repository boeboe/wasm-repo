package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMFile struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	DownloadAlias string    `gorm:"type:text;not null"`
	Filename      string    `gorm:"type:text;not null"`
	Path          string    `gorm:"type:text;not null"`
	Sha256        string    `gorm:"type:varchar(255);not null"`
	Size          int       `gorm:"type:int;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ReleaseID     uuid.UUID `gorm:"type:uuid;not null"`
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (wf *WASMFile) BeforeCreate(tx *gorm.DB) (err error) {
	wf.ID = uuid.New() // Generate a new UUID for the record.
	return
}

func (wf *WASMFile) SetDownloadAlias(pluginName, releaseVersion string) {
	wf.DownloadAlias = fmt.Sprintf("%s-%s.wasm", pluginName, releaseVersion)
}
