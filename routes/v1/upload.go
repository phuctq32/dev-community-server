package routesv1

import (
	"dev_community_server/components/appctx"
	"dev_community_server/modules/upload/transport/api"
	"github.com/gin-gonic/gin"
)

func NewUploadRoutes(appCtx appctx.AppContext, group *gin.RouterGroup) {
	uploadHandler := api.NewUploadHandler(appCtx)

	uploadRouter := group.Group("/upload")
	{
		uploadRouter.POST("/images/single", uploadHandler.UploadSingleImage())
		uploadRouter.POST("/images/multiple", uploadHandler.UploadMultipleImage())
	}
}
