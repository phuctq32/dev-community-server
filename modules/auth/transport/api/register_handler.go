package api

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userEntity.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := hdl.appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		if err := hdl.business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Register successfully! We sent a code to verify your email, please check!",
		})
	}
}
