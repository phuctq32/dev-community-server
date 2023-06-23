package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string               `bson:"name" json:"name"`
	Description    string               `bson:"description" json:"description"`
	ModeratorIds   []primitive.ObjectID `bson:"moderator_ids" json:"-"`
}

func NewTopic(data *TopicCreate) *Topic {
	return &Topic{
		Name:        data.Name,
		Description: data.Description,
	}
}
