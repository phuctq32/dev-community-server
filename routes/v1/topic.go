package routesv1

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/middlewares"
	"dev_community_server/modules/topic/transport/api"
	"github.com/gin-gonic/gin"
)

func NewTopicRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	topicHandler := api.NewTopicHandler(appCtx)

	topicRouter := group.Group("/topics")
	{
		topicRouter.GET("", topicHandler.GetTopics())

	}

	topicProtectedRouter := topicRouter.Use(middlewares.Authorize(appCtx))
	{
		topicProtectedRouter.POST("", middlewares.CheckRole([]common.RoleType{common.ADMINISTRATOR}), topicHandler.CreateTopic())
	}
}
