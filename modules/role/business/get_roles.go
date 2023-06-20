package business

import (
	"context"
	"dev_community_server/modules/role/entity"
)

func (biz *roleBusiness) GetRoles(ctx context.Context) ([]*entity.Role, error) {
	roles, err := biz.repo.Find(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
