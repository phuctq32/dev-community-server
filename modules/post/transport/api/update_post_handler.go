package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/entity"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (hdl *postHandler) UpdatePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.PostUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
		}

		// Get user id
		requester := c.MustGet(common.ReqUser).(common.Requester)
		data.Author = requester.(*userEntity.User)

		// Get post id
		postId := strings.TrimSpace(c.Param("id"))
		data.Id = &postId

		post, err := hdl.business.UpdatePost(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Updated post successfully", post))
	}
}
