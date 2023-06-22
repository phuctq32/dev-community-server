package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoId struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

type MongoTimestamps struct {
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
