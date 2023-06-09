package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *roleHandler) GetRoles(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := hdl.biz.GetRoles(c.Request.Context())
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", roles))
	}
}
