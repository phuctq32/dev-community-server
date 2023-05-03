package business

import (
	"context"
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
)

func (biz *authBusiness) Register(ctx context.Context, data *userEntity.UserCreate) error {
	existingUser, err := biz.repo.FindOne(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		if appErr, ok := err.(*common.AppError); ok {
			if appErr.Key != "NOT_FOUND" {
				return err
			}
		}
	}

	if existingUser != nil {
		return common.NewExistingError("user")
	}

	data.Password, err = biz.hash.HashPassword(data.Password)
	if err != nil {
		return common.NewServerError(err)
	}

	if err = biz.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
