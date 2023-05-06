package api

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func (hdl *userHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user userEntity.UserUpdate

		if err := c.ShouldBind(&user); err != nil {
			log.Println("loi bind")
			panic(common.NewServerError(err))
		}

		if err := hdl.appCtx.GetValidator().Validate(user); err != nil {
			panic(common.NewValidationError(err))
		}

		if err := hdl.business.UpdateUser(c.Request.Context(), strings.Trim(c.Param("id"), " "), &user); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Updated successfully!",
		})
	}
}