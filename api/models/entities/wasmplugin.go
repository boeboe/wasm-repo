package entities

import (
	"time"
)

type WASMPlugin struct {
	ID          string `gorm:"primaryKey;type:varchar(255)"`
	Name        string `gorm:"type:varchar(255);not null"`
	Owner       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Releases    []WASMRelease `gorm:"foreignKey:PluginID"`
}
