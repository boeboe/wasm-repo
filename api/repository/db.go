package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/boeboe/wasm-repo/api/repository/connectors"
	"gorm.io/gorm"
)

var (
	// Db is a global variable for the database connection
	db *gorm.DB
)

type Repositories struct {
	FileRepo    *WASMFileRepo
	PluginRepo  *WASMPluginRepo
	ReleaseRepo *WASMReleaseRepo
}

// ConnectDatabase initializes the database connection based on the DB_TYPE environment variable.
func ConnectDatabase() *Repositories {
	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		connectors.ConnectPostgres()
		db = connectors.PostgresDb
	case "mysql":
		connectors.ConnectMySQL()
		db = connectors.MySqlDb
	case "sqlite":
		connectors.ConnectSQLite()
		db = connectors.SqliteDb
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: postgres, mysql, sqlite.")
	}

	if db == nil {
		log.Fatalf("Failed to connect to the %s database", dbType)
	}
	log.Printf("Successfully connected to %s database.", dbType)

	fileRepo := &WASMFileRepo{Database: db}
	pluginRepo := &WASMPluginRepo{Database: db}
	releaseRepo := &WASMReleaseRepo{Database: db}
	return &Repositories{
		FileRepo:    fileRepo,
		PluginRepo:  pluginRepo,
		ReleaseRepo: releaseRepo,
	}
}

// handleDBError is a helper function to handle common database errors.
func handleDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("record not found: %v", err)
	}
	return err
}
