package models

import "github.com/boeboe/wasm-repo/api/models/sharedtypes"

type WASMRepository interface {
	CreateBinary(binary *sharedtypes.WASMBinary) error
	GetBinaryByID(id uint) (*sharedtypes.WASMBinary, error)
}
