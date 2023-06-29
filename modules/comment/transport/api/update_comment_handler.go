package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/comment/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *commentHandler) UpdateComment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.CommentUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		updatedCmt, err := hdl.business.UpdateComment(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Updated comment successfully", updatedCmt))
	}
}
