package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/hasher"
	"dev_community_server/components/jwt"
	"dev_community_server/modules/auth/business"
	"dev_community_server/modules/auth/entity"
	userEntity "dev_community_server/modules/user/entity"
	userRepo "dev_community_server/modules/user/repository"
)

type AuthBusiness interface {
	Register(ctx context.Context, data *userEntity.UserCreate) error
	Login(ctx context.Context, data *entity.UserLogin) (*string, error)
}

type authHandler struct {
	appCtx   appctx.AppContext
	business AuthBusiness
}

func NewAuthHandler(appCtx appctx.AppContext) *authHandler {
	authRepo := userRepo.NewUserRepository(appCtx.GetDbConnection())
	hashService := hasher.NewBcryptHash(12)
	jwtProvider := jwt.NewJwtProvider(appCtx.GetSecretKey())
	biz := business.NewAuthBusiness(authRepo, hashService, jwtProvider, 30*24*60)

	return &authHandler{appCtx: appCtx, business: biz}
}
