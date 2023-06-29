package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"time"
)

func (repo *userRepository) Create(ctx context.Context, user *userEntity.User) error {
	// Convert to object id
	roleOid, err := common.ToObjectId(user.RoleId)
	if err != nil {
		return err
	}

	insertData := &userEntity.UserInsert{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Password:      user.Password,
		Birthday:      user.Birthday,
		Avatar:        common.DefaultAvatarUrl,
		RoleId:        roleOid,
		VerifiedToken: user.VerifiedToken,
	}

	if _, err = repo.userColl.InsertOne(ctx, insertData); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
