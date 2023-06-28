package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson"
)

type Topic struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string   `bson:"name" json:"name"`
	Description    string   `bson:"description" json:"description"`
	ModeratorIds   []string `bson:"moderator_ids" json:"-"`
}

func (*Topic) CollectionName() string { return "topics" }

func (topic *Topic) MarshalBSON() ([]byte, error) {
	return bson.Marshal(topic)
}
