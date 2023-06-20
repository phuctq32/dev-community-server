package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/role/business"
	"dev_community_server/modules/role/entity"
	"dev_community_server/modules/role/repository"
)

type RoleBusiness interface {
	CreateRole(ctx context.Context, data *entity.RoleCreate) (*entity.Role, error)
	GetRoles(ctx context.Context) ([]*entity.Role, error)
}

type roleHandler struct {
	appCtx appctx.AppContext
	biz    RoleBusiness
}

func NewRoleHandler(appCtx appctx.AppContext) *roleHandler {
	repo := repository.NewRoleRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	biz := business.NewRoleBusiness(repo)

	return &roleHandler{appCtx: appCtx, biz: biz}
}
