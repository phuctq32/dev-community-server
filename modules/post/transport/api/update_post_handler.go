package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (hdl *postHandler) UpdatePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.PostUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
		}

		// Get user id
		requester := c.MustGet(common.ReqUser).(common.Requester)
		userId := requester.GetUserId()
		data.AuthorId = &userId

		// Get post id
		postId := strings.TrimSpace(c.Param("id"))
		data.Id = &postId

		if err := hdl.business.UpdatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Updated post successfully", nil))
	}
}
