package routesv1

import (
	"dev_community_server/common"
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
		postRouter.GET("/search", postHandler.SearchPosts(appCtx))
	}

	postProtectedRouter := postRouter.Use(middlewares.Authorize(appCtx))
	{
		postProtectedRouter.POST("", postHandler.CreatePost(appCtx))
		postProtectedRouter.PATCH("/:id", postHandler.UpdatePost(appCtx))
		postProtectedRouter.DELETE("/:id")
		postProtectedRouter.POST("/:id/up-vote", postHandler.UpVote(appCtx))
		postProtectedRouter.POST("/:id/down-vote", postHandler.DownVote(appCtx))
		postProtectedRouter.POST("/:id/view")
		postProtectedRouter.POST("/:id/approve", middlewares.RequireRoles(common.Administrator, common.Moderator), postHandler.ApprovePostHandler(appCtx))
		postProtectedRouter.POST("/:id/block", middlewares.RequireRoles(common.Administrator, common.Moderator), postHandler.BlockPostHandler(appCtx))
		postProtectedRouter.POST("/:id/unblock", middlewares.RequireRoles(common.Administrator, common.Moderator), postHandler.UnblockPostHandler(appCtx))
		postProtectedRouter.GET("/pending", middlewares.RequireRoles(common.Administrator, common.Moderator), postHandler.GetPendingPost(appCtx))
	}

	currentUserRouter := group.Group("/me", middlewares.Authorize(appCtx))
	{
		currentUserRouter.GET("/saved-posts", postHandler.GetCurrentUserSavedPosts(appCtx))
		currentUserRouter.POST("/saved-posts/:postId", postHandler.SavePost(appCtx))
		currentUserRouter.DELETE("/saved-posts/:postId", postHandler.RemovePostFromCurrentUserSavedPosts(appCtx))
		currentUserRouter.DELETE("/saved-posts", postHandler.RemoveAllPostsFromCurrentUserSavedPosts(appCtx))
		// Pending posts
		currentUserRouter.GET("/pending-posts", postHandler.GetCurrentUserPendingPosts(appCtx))
		// Approved posts (posted posts)
		currentUserRouter.GET("/approved-posts")
	}
}
