package models

import (
	"log"
	"os"

	"github.com/boeboe/wasm-repo/api/models/connectors"
	"gorm.io/gorm"
)

var (
	// Db is a global variable for the database connection
	Db *gorm.DB

	// The wasm repository exposed to handlers
	Repo *WASMRepository
)

// ConnectDatabase initializes the database connection based on the DB_TYPE environment variable.
func ConnectDatabase() {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		connectors.ConnectPostgres() // Assuming you named the function ConnectPostgres
		Db = connectors.PostgresDb
	case "mysql":
		connectors.ConnectMySQL() // Assuming you named the function ConnectMySQL
		Db = connectors.MySqlDb
	case "sqlite":
		connectors.ConnectSQLite() // Assuming you named the function ConnectSQLite
		Db = connectors.SqliteDb
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: postgres, mysql, sqlite.")
	}

	if Db == nil {
		log.Fatalf("Failed to connect to the %s database", dbType)
	}

	Repo = &WASMRepository{Database: Db}

	log.Printf("Successfully connected to %s database.", dbType)
}
