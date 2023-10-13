package boltrepo

import (
	"log"
	"time"

	"go.etcd.io/bbolt"
)

var BoltDb *bbolt.DB

// Connect initializes the bbolt file database connection.
func Connect() {
	log.Println("Connecting to bbolt file database...")

	db, err := bbolt.Open("build/wasm.db", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("Failed to open bbolt database: %v", err)
	}

	// Initialize buckets
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("WASMBinary"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("WASMMetadata"))
		return err
	})

	if err != nil {
		log.Fatalf("Failed to create buckets: %v", err)
	}

	BoltDb = db
	log.Println("Successfully connected to bbolt file database.")
}
