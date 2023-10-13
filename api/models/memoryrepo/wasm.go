package memoryrepo

import (
	"time"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"github.com/hashicorp/go-memdb"
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

func CreateMemDBSchema() (*memdb.DBSchema, error) {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"WASMBinaries": {
				Name: "WASMBinaries",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "name",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			"WASMMetadatas": {
				Name: "WASMMetadatas",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}, nil
}
