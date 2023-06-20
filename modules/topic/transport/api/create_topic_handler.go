package api

import (
	"dev_community_server/modules/topic/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *topicHandler) CreateTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.TopicCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := hdl.appCtx.GetValidator().Validate(&data); err != nil {
			panic(err)
			return
		}

		topic, err := hdl.biz.CreateTopic(c.Request.Context(), &data)
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Create topic successfully",
			"topic":   topic,
		})
	}
}
