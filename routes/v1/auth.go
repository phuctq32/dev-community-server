package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/modules/auth/transport/api"
	"github.com/gin-gonic/gin"
)

func NewAuthRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	authHandler := api.NewAuthHandler(appCtx)

	authRouter := group.Group("/auth")
	{
		authRouter.POST("/signup", authHandler.RegisterHandler())
		authRouter.POST("/login", authHandler.LoginHandler())
		authRouter.POST("/verification/:verifiedToken", authHandler.VerifyEmail())
	}
}
