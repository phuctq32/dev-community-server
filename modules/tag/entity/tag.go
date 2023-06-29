package entity

import (
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagInsert struct {
	common.MongoId `bson:",inline"`
	Name           string             `bson:"name"`
	TopicId        primitive.ObjectID `bson:"topic_id"`
}

type Tag struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string        `bson:"name" json:"name"`
	TopicId        string        `bson:"topic_id" json:"-"`
	Topic          *entity.Topic `bson:"-" json:"topic,omitempty"`
}

func (*Tag) CollectionName() string { return "tags" }
