package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

func (repo *topicRepository) Find(ctx context.Context, filter common.Filter) ([]entity.Topic, error) {
	opts := options.Find().SetLimit(int64(*filter.Limit)).SetSkip((int64(math.Abs(float64(*filter.Page-1))) * int64(*filter.Limit)))

	cursor, err := repo.topicColl.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var topics []entity.Topic
	if err := cursor.All(ctx, &topics); err != nil {
		return nil, err
	}

	return topics, nil
}

func (repo *topicRepository) Find1(ctx context.Context, filter map[string]interface{}) ([]entity.Topic, error) {
	if err := common.ConvertFieldToObjectId(filter, map[string]string{"moderator_ids": "moderator_ids"}); err != nil {
		return nil, common.NewServerError(err)
	}

	cursor, err := repo.topicColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var topics []entity.Topic
	if err := cursor.All(ctx, &topics); err != nil {
		return nil, err
	}

	if topics == nil {
		topics = []entity.Topic{}
	}

	return topics, nil
}
