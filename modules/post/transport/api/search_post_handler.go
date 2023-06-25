package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (hdl *postHandler) SearchPost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination *common.Pagination
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

			pagination = &common.Pagination{
				Limit: &intLimit,
				Page:  &intPage,
			}
		}

		searchTerm := ""
		if search, ok := c.GetQuery("q"); ok {
			searchTerm = strings.TrimSpace(search)
		}

		posts, paginationInfo, err := hdl.business.SearchPosts(c.Request.Context(), &searchTerm, pagination)
		if err != nil {
			panic(err)
		}

		postCount := len(posts)

		c.JSON(http.StatusOK, common.NewFullResponse("", posts, &postCount, paginationInfo))
	}
}
