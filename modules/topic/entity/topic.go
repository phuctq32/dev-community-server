package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopicInsert struct {
	common.MongoId         `bson:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Name                   string               `bson:"name"`
	Description            string               `bson:"description"`
	ModeratorIds           []primitive.ObjectID `bson:"moderator_ids"`
}

type Topic struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Name                   string   `bson:"name" json:"name"`
	Description            string   `bson:"description" json:"description"`
	ModeratorIds           []string `bson:"moderator_ids" json:"-"`
}

func (*Topic) CollectionName() string { return "topics" }
