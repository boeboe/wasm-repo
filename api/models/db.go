package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hashicorp/go-memdb"
	"go.etcd.io/bbolt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB is a global variable for the PostgreSQL database connection
	DB *gorm.DB
	// MongoDb is a global variable for the MongoDB database connection
	MongoDb *mongo.Database
	// MemDB is a global variable for the in-memory database connection
	MemDB *memdb.MemDB
	// BoltDB is a global variable for the bbolt file database connection
	BoltDB *bbolt.DB
)

// ConnectDatabase initializes the database connection based on the DB_TYPE environment variable.
func ConnectDatabase() {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mongo":
		connectMongoDB()
	case "postgres":
		connectPostgres()
	case "memdb":
		connectMemDB()
	case "bbolt":
		connectBoltDB()
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: mongo, postgres.")
	}
}

// connectPostgres initializes the PostgreSQL database connection with a timeout.
func connectPostgres() {
	log.Println("Connecting to PostgreSQL database...")

	// Adding a connection timeout parameter to the DSN
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Europe/Brussels connect_timeout=10"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the PostgreSQL database: %v", err)
	}

	err = db.AutoMigrate(&WASMBinary{}, &WASMMetadata{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	DB = db
	log.Println("Successfully connected to PostgreSQL database.")
}

// connectMongoDB initializes the MongoDB database connection.
func connectMongoDB() {
	log.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI("mongodb://gorm:gorm@localhost:27018")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	MongoDb = client.Database("gorm")
	log.Println("Successfully connected to MongoDB.")
}

// connectMemDB initializes the in-memory database connection.
func connectMemDB() {
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

	MemDB = db
	log.Println("Successfully initialized in-memory database.")
}

// connectBoltDB initializes the bbolt file database connection.
func connectBoltDB() {
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

	BoltDB = db
	log.Println("Successfully connected to bbolt file database.")
}
