package models

import (
	"log"
	"os"

	"github.com/boeboe/wasm-repo/api/models/connectors"
	"gorm.io/gorm"
)

var (
	// Db is a global variable for the database connection
	db *gorm.DB
)

// ConnectDatabase initializes the database connection based on the DB_TYPE environment variable.
func ConnectDatabase() *WASMRepository {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		connectors.ConnectPostgres() // Assuming you named the function ConnectPostgres
		db = connectors.PostgresDb
	case "mysql":
		connectors.ConnectMySQL() // Assuming you named the function ConnectMySQL
		db = connectors.MySqlDb
	case "sqlite":
		connectors.ConnectSQLite() // Assuming you named the function ConnectSQLite
		db = connectors.SqliteDb
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: postgres, mysql, sqlite.")
	}

	if db == nil {
		log.Fatalf("Failed to connect to the %s database", dbType)
	}

	repo := &WASMRepository{Database: db}

	log.Printf("Successfully connected to %s database.", dbType)
	return repo
}
