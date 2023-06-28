package entity

import (
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type Tag struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string        `bson:"name" json:"name"`
	TopicId        string        `bson:"topic_id" json:"-"`
	Topic          *entity.Topic `bson:"-" json:"topic,omitempty"`
}

func (*Tag) CollectionName() string { return "tags" }

func (tag *Tag) MarshalBSON() ([]byte, error) {
	dataBytes, _ := bson.Marshal(tag)

	var bm common.BsonMap
	if err := bson.Unmarshal(dataBytes, &bm); err != nil {
		return nil, err
	}

	if err := bm.ToObjectId("topic_id"); err != nil {
		return nil, err
	}

	return bson.Marshal(bm)
}
