package mongorepo

import (
	"context"
	"fmt"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Database *mongo.Database
}

// CreateBinary saves the binary data to the MongoDB.
func (r *MongoRepository) CreateBinary(binary *sharedtypes.WASMBinary) error {
	collection := r.Database.Collection("WASMBinary")

	// Convert the ID from uint to primitive.ObjectID
	binary.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.TODO(), binary)
	if err != nil {
		return err
	}

	return nil
}

// GetBinaryByID retrieves the binary data by ID from the MongoDB.
func (r *MongoRepository) GetBinaryByID(id uint) (*sharedtypes.WASMBinary, error) {
	collection := r.Database.Collection("WASMBinary")

	// Convert the ID from uint to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(fmt.Sprintf("%x", id))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	result := &sharedtypes.WASMBinary{}

	err = collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Binary with ID %d not found", id)
		}
		return nil, err
	}

	return result, nil
}
