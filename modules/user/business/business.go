package business

import (
	"context"
	"dev_community_server/modules/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.UserCreate) error
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.User, error)
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type userBusiness struct {
	repo UserRepository
}

func NewUserBusiness(repo UserRepository) *userBusiness {
	return &userBusiness{repo: repo}
}
