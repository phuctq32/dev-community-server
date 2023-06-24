package business

import (
	"context"
	"dev_community_server/modules/role/entity"
)

type RoleRepository interface {
	Create(ctx context.Context, data *entity.RoleCreate) (*entity.Role, error)
	Find(ctx context.Context) ([]entity.Role, error)
}

type roleBusiness struct {
	repo RoleRepository
}

func NewRoleBusiness(roleRepo RoleRepository) *roleBusiness {
	return &roleBusiness{repo: roleRepo}
}
