package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
)

func (biz *userBusiness) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	userFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(userFilter, "id", id); err != nil {
		return nil, err
	}
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	roleFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(roleFilter, "id", user.RoleId); err != nil {
		return nil, err
	}
	role, err := biz.roleRepo.FindOne(ctx, roleFilter)
	if err != nil {
		return nil, err
	}
	user.Role = role.Name

	return user, nil
}
