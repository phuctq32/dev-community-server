package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *topicHandler) GetTopicById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct {
			Id *string `json:"id" validate:"required,mongodb"`
		}
		topicId := c.Param("id")
		data.Id = &topicId
		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		topic, err := hdl.biz.GetTopicById(c.Request.Context(), topicId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", topic))
	}
}
