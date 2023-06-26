package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *topicHandler) GetTopics(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		topics, err := hdl.biz.GetTopics(c.Request.Context(), map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSimpleResponse("", topics))
	}
}
