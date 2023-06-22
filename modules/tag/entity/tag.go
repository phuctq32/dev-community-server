package entity

import (
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	TopicId primitive.ObjectID `bson:"topic_id" json:"-"`
	Topic   *entity.Topic      `bson:"-" json:"topic"`
}
