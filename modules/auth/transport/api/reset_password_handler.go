package api

import (
	"dev_community_server/common"
	"dev_community_server/modules/auth/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) RestPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.UserResetPassword
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := hdl.appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		resetToken := c.Param("reset_token")

		if err := hdl.business.ResetPassword(c.Request.Context(), resetToken, *data.Password); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Reset password successfully", nil))
	}
}
