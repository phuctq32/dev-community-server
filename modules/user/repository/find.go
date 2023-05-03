package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *userRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error) {
	var user userEntity.User

	if id, ok := filter["id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter["_id"] = objId
		delete(filter, "id")
	}

	if err := repo.userColl.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NewNotFoundError("user", err)
		}

		return nil, err
	}

	return &user, nil
}
