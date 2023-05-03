package business

import (
	"context"
	"dev_community_server/components/hasher"
	userEntity "dev_community_server/modules/user/entity"
)

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*userEntity.User, error)
	Create(ctx context.Context, data *userEntity.UserCreate) error
}

type authBusiness struct {
	repo AuthRepository
	hash hasher.PasswordHasher
}

func NewAuthBusiness(repo AuthRepository, hash hasher.PasswordHasher) *authBusiness {
	return &authBusiness{repo: repo, hash: hash}
}
