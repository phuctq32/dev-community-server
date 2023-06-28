package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (repo *topicRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Topic, error) {
	if err := common.BsonMap(data).ToListObjectId("moderator_ids"); err != nil {
		return nil, common.NewServerError(err)
	}
	data["updated_at"] = time.Now()

	var updatedTopic entity.Topic
	if err := repo.topicColl.FindOneAndUpdate(
		ctx,
		filter,
		data,
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&updatedTopic); err != nil {
		return nil, common.NewServerError(err)
	}

	return &updatedTopic, nil
}
