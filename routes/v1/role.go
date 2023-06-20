package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/modules/role/transport/api"
	"github.com/gin-gonic/gin"
)

func NewRoleRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	roleHandler := api.NewRoleHandler(appCtx)

	roleRouter := group.Group("/roles")
	{
		roleRouter.GET("", roleHandler.GetRoles())
		roleRouter.POST("", roleHandler.CreateRole())
	}
}
