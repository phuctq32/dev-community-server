package business

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/hasher"
	"dev_community_server/components/jwt"
	"dev_community_server/components/mailer"
	"dev_community_server/modules/role/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
	Create(ctx context.Context, data *userEntity.UserCreate) error
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type RoleRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Role, error)
}

type authBusiness struct {
	appCtx        appctx.AppContext
	userRepo      UserRepository
	roleRepo      RoleRepository
	hash          hasher.MyHash
	jwtProvider   jwt.TokenProvider
	jwtExpiry     int
	emailProvider mailer.EmailProvider
}

func NewAuthBusiness(
	appCtx appctx.AppContext,
	repo UserRepository,
	roleRepo RoleRepository,
	hash hasher.MyHash,
	tokenProvider jwt.TokenProvider,
	tokenExpiry int,
	emailProvider mailer.EmailProvider,
) *authBusiness {
	return &authBusiness{
		appCtx:        appCtx,
		userRepo:      repo,
		roleRepo:      roleRepo,
		hash:          hash,
		jwtProvider:   tokenProvider,
		jwtExpiry:     tokenExpiry,
		emailProvider: emailProvider,
	}
}
