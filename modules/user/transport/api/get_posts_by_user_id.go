package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) GetPostsByUserId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := hdl.business.GetPostsByUserId(c.Request.Context(), c.Param("userId"))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", posts))
	}
}
