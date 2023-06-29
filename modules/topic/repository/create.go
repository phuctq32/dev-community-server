package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo *topicRepository) Create(ctx context.Context, topic *entity.Topic) (*entity.Topic, error) {
	insertData := &entity.TopicInsert{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:         topic.Name,
		Description:  topic.Description,
		ModeratorIds: []primitive.ObjectID{},
	}
	result, err := repo.topicColl.InsertOne(ctx, insertData)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	topic.Id = &insertedId
	topic.CreatedAt = insertData.CreatedAt
	topic.UpdatedAt = insertData.UpdatedAt

	return topic, nil
}
