package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/hasher"
	"dev_community_server/modules/auth/business"
	userEntity "dev_community_server/modules/user/entity"
	userRepo "dev_community_server/modules/user/repository"
)

type AuthBusiness interface {
	Register(ctx context.Context, data *userEntity.UserCreate) error
}

type authHandler struct {
	appCtx   appctx.AppContext
	business AuthBusiness
}

func NewAuthHandler(appCtx appctx.AppContext) *authHandler {
	authRepo := userRepo.NewUserRepository(appCtx.GetDbConnection())
	hashService := hasher.NewBcryptHash(12)
	biz := business.NewAuthBusiness(authRepo, hashService)

	return &authHandler{appCtx: appCtx, business: biz}
}
