package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	"dev_community_server/modules/post/transport/api"
	"github.com/gin-gonic/gin"
)

func NewPostRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	postHandler := api.NewPostHandler(appCtx)

	postRouter := group.Group("/posts")
	{
		postRouter.GET("/:userId", postHandler.GetPostByUserId())
		postRouter.GET("/search", postHandler.SearchPost())
	}

	postProtectedRouter := postRouter.Use(middlewares.Authorize(appCtx))
	{
		postProtectedRouter.POST("", postHandler.CreatePost())
	}
}
