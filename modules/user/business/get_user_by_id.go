package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
)

func (biz *userBusiness) GetUserById(ctx context.Context, id string) (user *entity.User, err error) {
	user, err = biz.userRepo.FindOne(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	return user, nil
}
