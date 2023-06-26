package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (hdl *postHandler) GetPosts(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := common.DefaultPagination
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

			pagination.Limit = intLimit
			pagination.Page = intPage
		}

		posts, paginationInfo, err := hdl.business.GetPosts(c.Request.Context(), map[string]interface{}{}, pagination)
		if err != nil {
			panic(err)
		}
		postCount := len(posts)

		c.JSON(http.StatusOK, common.NewFullResponse("", posts, &postCount, paginationInfo))
	}
}
