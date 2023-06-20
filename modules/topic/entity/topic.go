package entity

import (
	"dev_community_server/common"
	"time"
)

type Topic struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Name               string `bson:"name" json:"name"`
	Description        string `bson:"description" json:"description"`
}

func NewTopic(data *TopicCreate) *Topic {
	now := time.Now()
	return &Topic{
		ModelCommon: common.ModelCommon{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Name:        data.Name,
		Description: data.Description,
	}
}
