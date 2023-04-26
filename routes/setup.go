package routes

import (
	"dev_community_server/components/appctx"
	routesv1 "dev_community_server/routes/v1"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(appCtx appctx.AppContext, router *gin.Engine) {
	v1 := router.Group("/v1")
	routesv1.UserRoutes(appCtx, v1)
}
