package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (hdl *postHandler) GetTrendingPosts(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		quantity := 5
		if n, ok := c.GetQuery("n"); ok {
			iQuantity, err := strconv.Atoi(n)
			if err != nil {
				panic(common.NewCustomBadRequestError("n must be integer greater than 0"))
			}
			quantity = iQuantity
		}

		posts, err := hdl.business.GetTrendingPosts(c.Request.Context(), quantity)
		if err != nil {
			panic(err)
		}
		postCount := len(posts)

		c.JSON(http.StatusOK, common.NewFullResponse("", posts, &postCount, nil))
	}
}
