package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *topicRepository) Create(ctx context.Context, data *entity.TopicCreate) (*entity.Topic, error) {
	topic := entity.NewTopic(data)

	result, err := repo.topicColl.InsertOne(ctx, topic)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	topic.Id = result.InsertedID.(primitive.ObjectID)

	return topic, nil
}
