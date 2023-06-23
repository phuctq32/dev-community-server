package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) VerifyEmail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiedToken := c.Param("verifiedToken")

		if err := hdl.business.VerifyEmail(c.Request.Context(), verifiedToken); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Your email is verified!", nil))
	}
}
