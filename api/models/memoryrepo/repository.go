package memoryrepo

import (
	"fmt"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"github.com/hashicorp/go-memdb"
)

type MemoryRepository struct {
	Database *memdb.MemDB
}

// CreateBinary saves the binary data to the in-memory database.
func (r *MemoryRepository) CreateBinary(binary *sharedtypes.WASMBinary) error {
	txn := r.Database.Txn(true) // true means the transaction is writable
	defer txn.Abort()           // defer abort to ensure any failed transaction is not applied

	if err := txn.Insert("WASMBinary", binary); err != nil {
		return err
	}

	txn.Commit() // Commit the transaction
	return nil
}

// GetBinaryByID retrieves the binary data by ID from the in-memory database.
func (r *MemoryRepository) GetBinaryByID(id uint) (*sharedtypes.WASMBinary, error) {
	txn := r.Database.Txn(false) // false means the transaction is read-only
	defer txn.Abort()            // always abort read-only transactions

	raw, err := txn.First("WASMBinary", "id", id)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, fmt.Errorf("binary with ID %d not found", id)
	}

	binary, ok := raw.(*sharedtypes.WASMBinary)
	if !ok {
		return nil, fmt.Errorf("failed to cast to WASMBinary type")
	}

	return binary, nil
}
