package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *commentRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Comment, error) {
	if err := common.BsonMap(data).ToListMongoId("up_votes"); err != nil {
		return nil, common.NewServerError(err)
	}
	if err := common.BsonMap(data).ToListMongoId("down_votes"); err != nil {
		return nil, common.NewServerError(err)
	}

	var updatedCmt entity.Comment
	if err := repo.commentColl.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{
			"$set": data,
			"$currentDate": bson.M{
				"updated_at": bson.M{"$type": "date"},
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&updatedCmt); err != nil {
		return nil, common.NewServerError(err)
	}

	return &updatedCmt, nil
}
