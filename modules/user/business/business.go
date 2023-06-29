package business

import (
	"context"
	"dev_community_server/components/hasher"
	entity2 "dev_community_server/modules/role/entity"
	"dev_community_server/modules/user/entity"
)

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.User, error)
	Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.User, error)
}

type RoleRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Role, error)
}

type userBusiness struct {
	userRepo UserRepository
	roleRepo RoleRepository
	hash     hasher.MyHash
}

func NewUserBusiness(userRepo UserRepository, roleRepo RoleRepository, hash hasher.MyHash) *userBusiness {
	return &userBusiness{userRepo: userRepo, roleRepo: roleRepo, hash: hash}
}
