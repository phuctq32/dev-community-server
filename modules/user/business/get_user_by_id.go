package business

import (
	"context"
	"dev_community_server/modules/user/entity"
)

func (biz *userBusiness) GetUserById(ctx context.Context, id string) (user *entity.User, err error) {
	user, err = biz.repo.FindOne(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
