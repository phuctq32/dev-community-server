package entity

import (
	"dev_community_server/common"
)

type Topic struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
}

func NewTopic(data *TopicCreate) *Topic {
	return &Topic{
		Name:        data.Name,
		Description: data.Description,
	}
}
