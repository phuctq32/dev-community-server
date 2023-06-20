package repository

import "go.mongodb.org/mongo-driver/mongo"

type topicRepository struct {
	topicColl *mongo.Collection
}

func NewTopicRepository(db *mongo.Database) *topicRepository {
	return &topicRepository{topicColl: db.Collection("topics")}
}
