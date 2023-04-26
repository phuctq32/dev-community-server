package routes

import (
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(appCtx *appctx.AppContext, router *gin.Engine) {
	v1 := router.Group("/v1")
}
