package mongorepo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database

// Connect initializes the MongoDB database connection.
func Connect() {
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
