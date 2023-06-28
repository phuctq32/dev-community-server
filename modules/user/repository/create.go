package repository

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"time"
)

func (repo *userRepository) Create(ctx context.Context, data *userEntity.UserCreate) error {
	birthday := time.Time(data.Birthday)
	if _, err := repo.userColl.InsertOne(ctx, &userEntity.User{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Password:  data.Password,
		Birthday:  &birthday,
		Avatar:    common.DefaultAvatarUrl,
		RoleId:    data.RoleId,
		VerifiedToken: &userEntity.Token{
			Token:     data.VerifiedToken,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour * 24 * 7)),
		},
		IsVerified:   false,
		SavedPostIds: []string{},
	}); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
