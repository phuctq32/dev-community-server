package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *userRepository) Update(ctx context.Context, id string, data map[string]interface{}) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewServerError(err)
	}

	var existingUser userEntity.User

	if err = repo.userColl.FindOne(ctx, bson.M{"_id": objId}).Decode(&existingUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return common.NewNotFoundError("user", err)
		}

		return err
	}

	if _, err = repo.userColl.UpdateByID(ctx, objId, bson.M{
		"$set": data,
		"$currentDate": bson.M{
			"updated_at": bson.M{"$type": "date"},
		},
	}); err != nil {
		return err
	}

	return nil
}
