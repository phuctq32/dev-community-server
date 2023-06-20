package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (hdl *topicHandler) GetTopics() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter common.Filter
		limit, gotLimit := c.GetQuery("limit")
		page, gotPage := c.GetQuery("page")
		if gotLimit && gotPage {
			intPage, err := strconv.Atoi(page)
			if err != nil {
				panic(err)
			}
			intLimit, err := strconv.Atoi(limit)
			if err != nil {
				panic(err)
			}

			if intPage < 1 {
				intPage = 1
			}

			if intLimit < 10 {
				intLimit = 10
			}

			filter.Page = &intPage
			filter.Limit = &intLimit
		} else {
			var (
				intPage  int = 1
				intLimit int = 10
			)
			filter.Page = &intPage
			filter.Limit = &intLimit
		}

		topics, err := hdl.biz.GetTopics(c.Request.Context(), filter)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"topics": topics,
			"page":   filter.Page,
			"count":  len(topics),
		})
	}
}
