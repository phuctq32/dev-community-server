package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	"dev_community_server/modules/tag/transport/api"
	"github.com/gin-gonic/gin"
)

func NewTagRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	tagHandler := api.NewTagHandler(appCtx)

	tagRouter := group.Group("/tags")
	{

	}

	tagProtectedRouter := tagRouter.Use(middlewares.Authorize(appCtx))
	{
		tagProtectedRouter.POST("", tagHandler.CreateTag())
	}
}
