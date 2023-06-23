package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (hdl *tagHandler) GetTagsByTopicId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data struct {
			TopicId string `json:"topicId" validate:"required,mongodb"`
		}
		topicId := c.Query("topicId")
		log.Println(topicId)
		data.TopicId = topicId

		if err := hdl.appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		tags, err := hdl.biz.GetTagsByTopicId(c.Request.Context(), topicId)
		length := len(tags)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewFullResponse("", tags, &length, nil, nil))
	}
}
