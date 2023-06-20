package api

import (
	"dev_community_server/modules/role/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *roleHandler) CreateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.RoleCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := hdl.appCtx.GetValidator().Validate(&data); err != nil {
			panic(err)
			return
		}

		role, err := hdl.biz.CreateRole(c.Request.Context(), &data)
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Create role successfully",
			"role":    role,
		})
	}
}
