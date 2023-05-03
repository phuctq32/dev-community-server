package business

import (
	"context"
	userEntity "dev_community_server/modules/user/entity"
)

type AuthRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
	Create(ctx context.Context, data *userEntity.UserCreate) error
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}

type authBusiness struct {
	repo AuthRepository
	hash PasswordHasher
}

func NewAuthBusiness(repo AuthRepository, hash PasswordHasher) *authBusiness {
	return &authBusiness{repo: repo, hash: hash}
}
