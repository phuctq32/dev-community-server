package business

import (
	"context"
	"dev_community_server/modules/role/entity"
)

func (biz *roleBusiness) CreateRole(ctx context.Context, data *entity.RoleCreate) (*entity.Role, error) {
	role, err := biz.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return role, nil
}
