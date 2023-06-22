package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.ReqUser).(common.Requester)

		user, err := hdl.business.GetUserById(c.Request.Context(), requester.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", user))
	}
}
