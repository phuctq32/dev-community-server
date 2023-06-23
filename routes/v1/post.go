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
		postRouter.GET("", postHandler.GetPosts(appCtx))
		postRouter.GET("/:id", postHandler.GetPostById(appCtx))
		postRouter.GET("/search", postHandler.SearchPost(appCtx))
	}

	postProtectedRouter := postRouter.Use(middlewares.Authorize(appCtx))
	{
		postProtectedRouter.POST("", postHandler.CreatePost(appCtx))
		postProtectedRouter.PATCH("/:id", postHandler.UpdatePost(appCtx))
	}
}
