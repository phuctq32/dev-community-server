package repository

import (
	"context"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *topicRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Topic, error) {
	var topic entity.Topic
	if err := repo.topicColl.FindOne(ctx, filter).Decode(&topic); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &topic, nil
}
