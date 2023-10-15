package connectors

import (
	"log"

	"github.com/boeboe/wasm-repo/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

// Connect initializes the PostgreSQL database connection with a timeout.
func ConnectPostgres() {
	log.Println("Connecting to PostgreSQL database...")

	// Adding a connection timeout parameter to the DSN
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Europe/Brussels connect_timeout=10"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the PostgreSQL database: %v", err)
	}

	// Auto-migrate the new model structs
	err = db.AutoMigrate(&models.WASMPlugin{}, &models.WASMRelease{}, &models.WASMFile{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	PostgresDb = db
	log.Println("Successfully connected to PostgreSQL database.")
}
