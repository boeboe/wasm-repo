package models

import (
	"log"
	"os"

	"github.com/boeboe/wasm-repo/api/models/entities"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

	var err error

	switch dbType {
	case "postgres":
		Db, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	case "mysql":
		Db, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	case "sqlite":
		Db, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_DSN")), &gorm.Config{})
	default:
		log.Fatalf("Unknown DB_TYPE provided. Supported types are: postgres, mysql, sqlite.")
	}

	if err != nil {
		log.Fatalf("Failed to connect to the %s database: %v", dbType, err)
	}

	// AutoMigrate models
	if err := Db.AutoMigrate(&entities.WASMPlugin{}, &entities.WASMRelease{}, &entities.WASMLocation{}); err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	Repo = &WASMRepository{Database: Db}

	log.Printf("Successfully connected to %s database.", dbType)
}
