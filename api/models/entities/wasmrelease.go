package entities

import (
	"time"
)

type WASMRelease struct {
	ID          string `gorm:"primaryKey;type:varchar(255)"`
	Version     string `gorm:"type:varchar(255);not null"`
	Sha256      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Size        int    `gorm:"type:int;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Location    WASMLocation `gorm:"foreignKey:ReleaseID"`
	PluginID    string       `gorm:"type:varchar(255);not null"`
}

type WASMLocation struct {
	ID        string `gorm:"primaryKey;type:varchar(255)"`
	Type      string `gorm:"type:varchar(255);not null"`
	Location  string `gorm:"type:text;not null"`
	ReleaseID string `gorm:"type:varchar(255);not null"`
}
