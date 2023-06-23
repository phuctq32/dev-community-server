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
	biz RoleBusiness
}

func NewRoleHandler(appCtx appctx.AppContext) *roleHandler {
	repo := repository.NewRoleRepository(appCtx.GetMongoDBConnection())
	biz := business.NewRoleBusiness(repo)

	return &roleHandler{biz: biz}
}
