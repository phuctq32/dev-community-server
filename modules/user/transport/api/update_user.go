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
			panic(common.NewServerError(err))
		}
		log.Printf("%+v", user)

		if err := appCtx.GetValidator().Validate(user); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		updatedUser, err := hdl.business.UpdateUser(c.Request.Context(), requester.GetUserId(), &user)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Updated successfully", updatedUser))
	}
}
