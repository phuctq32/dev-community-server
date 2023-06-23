package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (hdl *postHandler) GetPostById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := strings.TrimSpace(c.Param("id"))
		post, err := hdl.business.GetPostById(c.Request.Context(), &postId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", post))
	}
}
