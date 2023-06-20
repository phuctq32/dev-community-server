package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
)

func (repo *userRepository) Create(ctx context.Context, data *userEntity.UserCreate) error {
	user := userEntity.NewUser(data)

	role, err := repo.roleRepo.FindOne(ctx, map[string]interface{}{"name": "Member"})
	if err != nil {
		return err
	}
	user.RoleId = role.Id

	if _, err := repo.userColl.InsertOne(ctx, &user); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
