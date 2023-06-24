package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *postRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Post, error) {
	if err := common.ConvertFieldToObjectId(filter, map[string]string{"id": "_id"}); err != nil {
		return nil, err
	}
	if err := common.ConvertFieldToObjectId(data, map[string]string{"topic_id": "topic_id", "tag_ids": "tag_ids"}); err != nil {
		return nil, err
	}

	var post entity.Post
	opts := options.FindOneAndUpdate().SetReturnDocument(1)
	if err := repo.postColl.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": data,
		"$currentDate": bson.M{
			"updated_at": bson.M{"$type": "date"},
		},
	}, opts).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}
