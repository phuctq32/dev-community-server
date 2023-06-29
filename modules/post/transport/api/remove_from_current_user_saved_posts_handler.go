package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *postHandler) RemovePostFromCurrentUserSavedPosts(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct {
			PostId string `json:"postId" validate:"required,mongodb"`
		}{}
		params.PostId = c.Param("postId")

		if err := appCtx.GetValidator().Validate(&params); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		posts, err := hdl.business.RemovePostFromSavedPosts(c.Request.Context(), params.PostId, requester.GetUserId())
		if err != nil {
			panic(err)
		}
		postCount := len(posts)

		c.JSON(http.StatusOK, common.NewFullResponse("Removed post from saved list", posts, &postCount, nil))
	}
}
