package api

import (
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) ChangePasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.UserUpdate
		if err := c.ShouldBind(&user); err != nil {
			panic(err)
		}

		if err := hdl.appCtx.GetValidator().Validate(user); err != nil {
			panic(common.NewValidationError(err))
		}

		if err := hdl.business.ChangePassword(c.Request.Context(), c.Param("id"), *user.Password); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Change passsword successfully",
		})
	}
}
