package repository

import (
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type topicRepository struct {
	topicColl *mongo.Collection
}

func NewTopicRepository(db *mongo.Database) *topicRepository {
	return &topicRepository{topicColl: db.Collection(new(entity.Topic).CollectionName())}
}
