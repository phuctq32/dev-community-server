package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) ForgotPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Email string `json:"email"`
		}

		if err := c.ShouldBind(&user); err != nil {
			panic(err)
		}

		if err := hdl.business.ForgotPassword(c.Request.Context(), user.Email); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Please check your email to reset password", nil))
	}
}
