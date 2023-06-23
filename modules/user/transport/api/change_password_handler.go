package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) ChangePasswordHandler(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.UserChangePassword
		if err := c.ShouldBind(&user); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(user); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		if err := hdl.business.ChangePassword(c.Request.Context(), requester.GetUserId(), &user); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Change password successfully", nil))
	}
}
