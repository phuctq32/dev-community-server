package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/hasher"
	repository2 "dev_community_server/modules/role/repository"
	"dev_community_server/modules/user/business"
	"dev_community_server/modules/user/entity"
	"dev_community_server/modules/user/repository"
)

type UserBusiness interface {
	GetUserById(ctx context.Context, id string) (user *entity.User, err error)
	UpdateUser(ctx context.Context, id string, data *entity.UserUpdate) (*entity.User, error)
	ChangePassword(ctx context.Context, id string, user *entity.UserChangePassword) error
}

type userHandler struct {
	business UserBusiness
}

func NewUserHandler(appCtx appctx.AppContext) *userHandler {
	userRepo := repository.NewUserRepository(appCtx.GetMongoDBConnection())
	roleRepo := repository2.NewRoleRepository(appCtx.GetMongoDBConnection())
	hash := hasher.NewBcryptHash(10)

	biz := business.NewUserBusiness(userRepo, roleRepo, hash)
	return &userHandler{business: biz}
}
