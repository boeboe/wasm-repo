package boltrepo

import (
	"encoding/json"
	"time"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"go.etcd.io/bbolt"
)

type WASMBinary struct {
	sharedtypes.WASMBinary

	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Binary    []byte       `json:"binary"`
	Metadata  WASMMetadata `json:"metadata"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type WASMMetadata struct {
	sharedtypes.WASMMetadata

	ID           string    `json:"id"`
	WASMBinaryID string    `json:"wasm_binary_id"`
	Description  string    `json:"description"`
	Owner        string    `json:"owner"`
	Version      string    `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func SaveWASMBinary(db *bbolt.DB, binary *WASMBinary) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("WASMBinaries"))
		if err != nil {
			return err
		}
		encodedBinary, err := json.Marshal(binary)
		if err != nil {
			return err
		}
		return b.Put([]byte(binary.ID), encodedBinary)
	})
}

func SaveWASMMetadata(db *bbolt.DB, metadata *WASMMetadata) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("WASMMetadatas"))
		if err != nil {
			return err
		}
		encodedMetadata, err := json.Marshal(metadata)
		if err != nil {
			return err
		}
		return b.Put([]byte(metadata.ID), encodedMetadata)
	})
}
