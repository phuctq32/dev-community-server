package business

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/jwt"
	"dev_community_server/components/mailer"
	userEntity "dev_community_server/modules/user/entity"
)

type AuthRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
	Create(ctx context.Context, data *userEntity.UserCreate) error
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}

type authBusiness struct {
	appCtx        appctx.AppContext
	repo          AuthRepository
	hash          PasswordHasher
	jwtProvider   jwt.TokenProvider
	jwtExpiry     int
	emailProvider mailer.EmailProvider
}

func NewAuthBusiness(
	appCtx appctx.AppContext,
	repo AuthRepository,
	hash PasswordHasher,
	tokenProvider jwt.TokenProvider,
	tokenExpiry int,
	emailProvider mailer.EmailProvider,
) *authBusiness {
	return &authBusiness{
		appCtx:        appCtx,
		repo:          repo,
		hash:          hash,
		jwtProvider:   tokenProvider,
		jwtExpiry:     tokenExpiry,
		emailProvider: emailProvider,
	}
}
