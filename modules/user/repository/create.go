package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
)

func (repo *userRepository) Create(ctx context.Context, data *userEntity.UserCreate) error {
	user := userEntity.NewUser(data)

	if _, err := repo.userColl.InsertOne(ctx, &user); err != nil {
		return common.NewServerError(err)
	}

	return nil
}