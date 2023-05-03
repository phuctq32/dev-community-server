package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StatusType uint8

const (
	DELETED StatusType = 0
	ACTIVE             = 1
)

type ModelCommon struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Status    StatusType         `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
