package business

import (
	"context"
	"dev_community_server/components/hasher"
	"dev_community_server/modules/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.UserCreate) error
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.User, error)
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type userBusiness struct {
	userRepo UserRepository
	hash     hasher.MyHash
}

func NewUserBusiness(repo UserRepository, hash hasher.MyHash) *userBusiness {
	return &userBusiness{userRepo: repo, hash: hash}
}
