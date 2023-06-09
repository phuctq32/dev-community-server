package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/topic/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *topicHandler) CreateTopic(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.TopicCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(err)
			return
		}

		topic, err := hdl.biz.CreateTopic(c.Request.Context(), &data)
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Create topic successfully", topic))
	}
}
