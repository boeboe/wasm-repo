package boltrepo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"go.etcd.io/bbolt"
)

type BoltRepository struct {
	Database *bbolt.DB
}

// CreateBinary saves the binary data to the bbolt database.
func (r *BoltRepository) CreateBinary(binary *sharedtypes.WASMBinary) error {
	return r.Database.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("WASMBinary"))
		if err != nil {
			return err
		}

		data, err := json.Marshal(binary)
		if err != nil {
			return err
		}

		return b.Put([]byte(fmt.Sprintf("%v", binary.ID)), data)
	})
}

// GetBinaryByID retrieves the binary data by ID from the bbolt database.
func (r *BoltRepository) GetBinaryByID(id uint) (*sharedtypes.WASMBinary, error) {
	var binary sharedtypes.WASMBinary

	err := r.Database.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("WASMBinary"))
		if b == nil {
			return errors.New("Bucket WASMBinary not found!")
		}

		data := b.Get([]byte(fmt.Sprintf("%v", id)))
		if data == nil {
			return errors.New("Binary not found!")
		}

		return json.Unmarshal(data, &binary)
	})

	if err != nil {
		return nil, err
	}

	return &binary, nil
}
