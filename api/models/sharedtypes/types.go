package sharedtypes

import (
	"time"
)

type WASMBinary struct {
	ID        interface{}
	Name      string
	Binary    []byte
	Metadata  WASMMetadata
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WASMMetadata struct {
	ID           interface{}
	WASMBinaryID interface{}
	Description  string
	Version      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
