package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/user/business"
	"dev_community_server/modules/user/entity"
	"dev_community_server/modules/user/repository"
)

type UserBusiness interface {
	GetUserById(ctx context.Context, id string) (user *entity.User, err error)
	UpdateUser(ctx context.Context, id string, data *entity.UserUpdate) error
}

type userHandler struct {
	appCtx   appctx.AppContext
	business UserBusiness
}

func NewUserHandler(appCtx appctx.AppContext) *userHandler {
	repo := repository.NewUserRepository(appCtx.GetDbConnection())

	biz := business.NewUserBusiness(repo)
	return &userHandler{appCtx: appCtx, business: biz}
}
