package routesv1

import (
	"dev_community_server/components/appctx"
	ginuser "dev_community_server/modules/user/transport/api"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(appCtx appctx.AppContext, v *gin.RouterGroup) {
	userHandler := ginuser.NewUserHandler(appCtx)

	userRouter := v.Group("/users")
	{
		userRouter.GET("/:id", userHandler.GetUserById())
		userRouter.PATCH("/:id", userHandler.UpdateUser())
	}
}
