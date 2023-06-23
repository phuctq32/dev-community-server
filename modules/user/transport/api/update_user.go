package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (hdl *userHandler) UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user userEntity.UserUpdate

		if err := c.ShouldBind(&user); err != nil {
			log.Println("loi bind")
			panic(common.NewServerError(err))
		}

		if err := appCtx.GetValidator().Validate(user); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		if err := hdl.business.UpdateUser(c.Request.Context(), requester.GetUserId(), &user); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Updated successfully", nil))
	}
}
