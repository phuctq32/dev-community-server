package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *userRepository) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.User, error) {
	if err := common.BsonMap(data).ToObjectId("role_id"); err != nil {
		return nil, common.NewServerError(err)
	}
	if err := common.BsonMap(data).ToListObjectId("saved_post_ids"); err != nil {
		return nil, common.NewServerError(err)
	}

	var updatedUser entity.User
	if err := repo.userColl.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": data,
		"$currentDate": bson.M{
			"updated_at": bson.M{"$type": "date"},
		},
	}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&updatedUser); err != nil {
		return nil, common.NewServerError(err)
	}

	return &updatedUser, nil
}
