package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *topicRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Topic, error) {
	if err := common.BsonMap(data).ToListObjectId("moderator_ids"); err != nil {
		return nil, common.NewServerError(err)
	}

	var updatedTopic entity.Topic
	if err := repo.topicColl.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{
			"$set": data,
			"$currentDate": bson.M{
				"updated_at": bson.M{"$type": "date"},
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&updatedTopic); err != nil {
		return nil, common.NewServerError(err)
	}

	return &updatedTopic, nil
}
