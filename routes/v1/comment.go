package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	"dev_community_server/modules/comment/transport/api"
	"github.com/gin-gonic/gin"
)

func NewCommentRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	commentHandler := api.NewCommentHandler(appCtx)

	commentRouter := group.Group("/comments")
	{
		commentRouter.GET("/:id", commentHandler.GetCommentById(appCtx))
		commentRouter.GET("/:id/replies", commentHandler.GetReplies(appCtx))
	}

	commentProtectedRouter := commentRouter.Use(middlewares.Authorize(appCtx))
	{
		commentProtectedRouter.POST("", commentHandler.CreateComment(appCtx))
		commentProtectedRouter.PATCH("/:id", commentHandler.UpdateComment(appCtx))
		commentProtectedRouter.POST("/:id/up-vote", commentHandler.UpVote(appCtx))
		commentProtectedRouter.POST("/:id/down-vote", commentHandler.DownVote(appCtx))
		commentProtectedRouter.POST("/:id/approve", commentHandler.ApproveComment(appCtx))
		commentProtectedRouter.POST("/:id/un-approve", commentHandler.UnApproveComment(appCtx))
	}
}
