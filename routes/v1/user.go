package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	ginuser "dev_community_server/modules/user/transport/api"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(appCtx appctx.AppContext, v *gin.RouterGroup) {
	userHandler := ginuser.NewUserHandler(appCtx)

	userRouter := v.Group("/users")
	{
		userRouter.GET("/:id", userHandler.GetUserById())
		userRouter.GET("/posts/:userId", userHandler.GetPostsByUserId())
	}
	userProtectedRouter := userRouter.Use(middlewares.Authorize(appCtx))
	{
		userProtectedRouter.GET("/profile", userHandler.GetProfile())
		userProtectedRouter.PATCH("/update", userHandler.UpdateUser())
		userProtectedRouter.PATCH("/change_password", userHandler.ChangePasswordHandler())
	}
}
