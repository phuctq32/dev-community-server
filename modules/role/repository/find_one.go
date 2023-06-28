package repository

import (
	"context"
	"dev_community_server/modules/role/entity"
)

func (repo *roleRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Role, error) {
	var role entity.Role
	if err := repo.roleColl.FindOne(ctx, filter).Decode(&role); err != nil {
		return nil, err
	}

	return &role, nil
}
