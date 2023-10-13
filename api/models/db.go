package models

import (
	"log"
	"os"

	"github.com/boeboe/wasm-repo/api/models/boltrepo"
	"github.com/boeboe/wasm-repo/api/models/memoryrepo"
	"github.com/boeboe/wasm-repo/api/models/mongorepo"
	"github.com/boeboe/wasm-repo/api/models/postgresrepo"

	"github.com/hashicorp/go-memdb"
	"go.etcd.io/bbolt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	// PostgresDb is a global variable for the PostgreSQL database connection
	PostgresDb *gorm.DB
	// MongoDb is a global variable for the MongoDB database connection
	MongoDb *mongo.Database
	// MemoryDb is a global variable for the in-memory database connection
	MemoryDb *memdb.MemDB
	// BoltDb is a global variable for the bbolt file database connection
	BoltDb *bbolt.DB

	// The wasm repository exposed to handlers
	Repo WASMRepository
)

// ConnectDatabase initializes the database connection based on the DB_TYPE environment variable.
func ConnectDatabase() {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		postgresrepo.Connect()
		Repo = &postgresrepo.PostgresRepository{Database: PostgresDb}
	case "mongo":
		mongorepo.Connect()
		Repo = &mongorepo.MongoRepository{Database: MongoDb}
	case "memdb":
		memoryrepo.Connect()
		Repo = &memoryrepo.MemoryRepository{Database: MemoryDb}
	case "bbolt":
		boltrepo.Connect()
		Repo = &boltrepo.BoltRepository{Database: BoltDb}
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: mongo, postgres.")
	}
}
