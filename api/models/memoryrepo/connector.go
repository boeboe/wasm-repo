package memoryrepo

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

var MemoryDb *memdb.MemDB

// Connect initializes the in-memory database connection.
func Connect() {
	log.Println("Initializing in-memory database...")

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"wasm": {
				Name: "wasm",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("Failed to create new in-memory database: %v", err)
	}

	MemoryDb = db
	log.Println("Successfully initialized in-memory database.")
}
