package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WASMPluginType string

const (
	HttpFilter    WASMPluginType = "HttpFilter"
	NetworkFilter WASMPluginType = "NetworkFilter"
	WasmService   WASMPluginType = "WasmService"
)

type WASMPlugin struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Description string    `gorm:"type:text"`
	Name        string    `gorm:"type:varchar(255);not null;unique"`
	Owner       string    `gorm:"type:varchar(255);not null"`
	Type        string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Releases    []WASMRelease `gorm:"foreignKey:PluginID"`
}

// BeforeCreate is a GORM hook that gets triggered before a new record is inserted into the database.
func (plugin *WASMPlugin) BeforeCreate(tx *gorm.DB) (err error) {
	plugin.ID = uuid.New() // Generate a new UUID for the record.
	return
}

type WASMPluginFilter struct {
	Name  string
	Owner string
	Type  string
}
