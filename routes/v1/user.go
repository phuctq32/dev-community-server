package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	ginuser "dev_community_server/modules/user/transport/api"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(appCtx appctx.AppContext, v *gin.RouterGroup) {
	userHandler := ginuser.NewUserHandler(appCtx)

	currentUserRouter := v.Group("/me")
	currentUserRouter.Use(middlewares.Authorize(appCtx))
	{
		currentUserRouter.GET("", userHandler.GetProfile(appCtx))
		currentUserRouter.PATCH("", userHandler.UpdateUser(appCtx))
		currentUserRouter.PATCH("/change-password", userHandler.ChangePasswordHandler(appCtx))
	}

	userRouter := v.Group("/users")
	{
		userRouter.GET("/:id", userHandler.GetUserById(appCtx))
	}
}
