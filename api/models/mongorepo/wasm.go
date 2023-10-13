package mongorepo

import (
	"time"

	"github.com/boeboe/wasm-repo/api/models/sharedtypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WASMBinary struct {
	sharedtypes.WASMBinary

	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Binary    []byte             `bson:"binary"`
	Metadata  WASMMetadata       `bson:"metadata"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type WASMMetadata struct {
	sharedtypes.WASMMetadata

	ID           primitive.ObjectID `bson:"_id,omitempty"`
	WASMBinaryID primitive.ObjectID `bson:"wasm_binary_id"`
	Description  string             `bson:"description"`
	Owner        string             `bson:"owner"`
	Version      string             `bson:"version"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
