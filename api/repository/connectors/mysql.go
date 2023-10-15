package connectors

import (
	"log"

	"github.com/boeboe/wasm-repo/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDb *gorm.DB

// Connect initializes the MySQL database connection with a timeout.
func ConnectMySQL() {
	log.Println("Connecting to MySQL database...")

	// DSN for MySQL
	dsn := "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the MySQL database: %v", err)
	}

	// Auto-migrate the model structs
	err = db.AutoMigrate(&models.WASMPlugin{}, &models.WASMRelease{}, &models.WASMFile{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}

	MySqlDb = db
	log.Println("Successfully connected to MySQL database.")
}
