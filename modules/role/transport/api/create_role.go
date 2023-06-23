package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/role/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *roleHandler) CreateRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.RoleCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(err)
			return
		}

		role, err := hdl.biz.CreateRole(c.Request.Context(), &data)
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Create role successfully", role))
	}
}
