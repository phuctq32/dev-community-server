package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *postRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Post, error) {
	if err := common.BsonMap(data).ToMongoId("topic_id"); err != nil {
		return nil, common.NewServerError(err)
	}
	if err := common.BsonMap(data).ToListMongoId("tag_ids"); err != nil {
		return nil, common.NewServerError(err)
	}
	if err := common.BsonMap(data).ToListMongoId("up_votes"); err != nil {
		return nil, common.NewServerError(err)
	}
	if err := common.BsonMap(data).ToListMongoId("down_votes"); err != nil {
		return nil, common.NewServerError(err)
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
