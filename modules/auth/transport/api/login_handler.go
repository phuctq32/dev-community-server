package api

import (
	"dev_community_server/common"
	"dev_community_server/modules/auth/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := hdl.appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		token, user, err := hdl.business.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", map[string]interface{}{
			"token": token,
			"user":  user,
		}))
	}
}
