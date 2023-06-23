package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/tag/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *tagHandler) CreateTag(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.TagCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		tag, err := hdl.biz.CreateTag(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Create tag successfully", tag))
	}
}
