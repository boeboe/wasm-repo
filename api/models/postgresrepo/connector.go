package postgresrepo

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

// Connect initializes the PostgreSQL database connection with a timeout.
func Connect() {
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

	PostgresDb = db
	log.Println("Successfully connected to PostgreSQL database.")
}
