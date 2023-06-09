package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *postHandler) CreatePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.PostCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)
		data.AuthorId = requester.GetUserId()

		post, err := hdl.business.CreatePost(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Created post successfully", post))
	}
}
