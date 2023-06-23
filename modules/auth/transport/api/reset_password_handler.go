package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/auth/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) RestPasswordHandler(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.UserResetPassword
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		resetToken := c.Param("resetToken")

		if err := hdl.business.ResetPassword(c.Request.Context(), resetToken, *data.Password); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Reset password successfully", nil))
	}
}
