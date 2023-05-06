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
		authRouter.POST("/verification/:verified_token", authHandler.VerifyEmail())
		authRouter.POST("/forgot_password", authHandler.ForgotPasswordHandler())
		authRouter.PATCH("/reset_password/:reset_token", authHandler.RestPasswordHandler())
	}
}
