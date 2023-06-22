package api

import (
	"dev_community_server/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := hdl.business.GetUserById(c.Request.Context(), c.Param("id"))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", user))
	}
}
