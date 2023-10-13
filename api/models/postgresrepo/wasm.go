package postgresrepo

import (
	"time"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
)

type WASMBinary struct {
	sharedtypes.WASMBinary

	ID        uint         `gorm:"primaryKey"`
	Name      string       `gorm:"type:varchar(255);not null;unique"`
	Binary    []byte       `gorm:"type:bytea"`
	Metadata  WASMMetadata `gorm:"foreignKey:WASMBinaryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WASMMetadata struct {
	sharedtypes.WASMMetadata

	ID           uint   `gorm:"primaryKey"`
	WASMBinaryID uint   `gorm:"not null"`
	Description  string `gorm:"type:varchar(255)"`
	Owner        string `gorm:"type:varchar(50)"`
	Version      string `gorm:"type:varchar(50)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
