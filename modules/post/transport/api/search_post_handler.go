package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (hdl *postHandler) SearchPost() gin.HandlerFunc {
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
				page  int = 1
				limit int = 10
			)
			filter.Page = &page
			filter.Limit = &limit
		}

		if search, ok := c.GetQuery("q"); ok {
			filter.Search = &search
		}

		posts, err := hdl.business.GetPosts(c.Request.Context(), filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	}
}
