package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) RegisterHandler(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userEntity.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		if err := hdl.business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Register successfully! We sent a code to verify your email, please check!", nil))
	}
}
