package routes

import (
	"dev_community_server/components/appctx"
	routesv1 "dev_community_server/routes/v1"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(appCtx appctx.AppContext, router *gin.Engine) {
	v1 := router.Group("/api/v1")
	routesv1.NewAuthRoutes(appCtx, v1)
	routesv1.NewUserRoutes(appCtx, v1)
	routesv1.NewUploadRoutes(appCtx, v1)
	routesv1.NewPostRoutes(appCtx, v1)
	routesv1.NewCommentRoutes(appCtx, v1)
	routesv1.NewRoleRoutes(appCtx, v1)
}
