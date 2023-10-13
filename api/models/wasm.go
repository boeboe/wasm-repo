package models

import (
	"time"
)

type WASMBinary struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `gorm:"type:varchar(255);not null;unique"`
	Binary    []byte       `gorm:"type:bytea"`
	Metadata  WASMMetadata `gorm:"foreignKey:WASMBinaryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WASMMetadata struct {
	ID           uint   `gorm:"primaryKey"`
	WASMBinaryID uint   `gorm:"not null"`
	Description  string `gorm:"type:varchar(255)"`
	Version      string `gorm:"type:varchar(50)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
