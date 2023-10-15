package connectors

import (
	"log"

	"github.com/boeboe/wasm-repo/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqliteDb *gorm.DB

// Connect initializes the SQLite database connection.
func ConnectSQLite() {
	log.Println("Connecting to SQLite database...")

	// DSN for SQLite
	dsn := "gorm.db"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}

	// Auto-migrate the model structs
	err = db.AutoMigrate(&models.WASMPlugin{}, &models.WASMRelease{}, &models.WASMFile{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	SqliteDb = db
	log.Println("Successfully connected to SQLite database.")
}
