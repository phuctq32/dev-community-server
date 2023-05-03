package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *userRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error) {
	var user userEntity.User

	if err := repo.userColl.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NewNotFoundError("user", err)
		}

		return nil, err
	}

	return &user, nil
}
