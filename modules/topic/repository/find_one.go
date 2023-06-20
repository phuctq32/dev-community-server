package repository

import (
	"context"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *topicRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Topic, error) {
	if id, ok := filter["id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter["_id"] = objId
		delete(filter, "id")
	}

	var topic entity.Topic
	if err := repo.topicColl.FindOne(ctx, filter).Decode(&topic); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &topic, nil
}
