package business

import (
	"context"
	"dev_community_server/components/jwt"
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
	repo          AuthRepository
	hash          PasswordHasher
	tokenProvider jwt.TokenProvider
	tokenExpiry   int
}

func NewAuthBusiness(
	repo AuthRepository,
	hash PasswordHasher,
	tokenProvider jwt.TokenProvider,
	tokenExpiry int,
) *authBusiness {
	return &authBusiness{
		repo:          repo,
		hash:          hash,
		tokenProvider: tokenProvider,
		tokenExpiry:   tokenExpiry,
	}
}
