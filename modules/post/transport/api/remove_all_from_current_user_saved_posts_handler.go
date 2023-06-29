package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *postHandler) RemoveAllPostsFromCurrentUserSavedPosts(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.ReqUser).(common.Requester)

		posts, err := hdl.business.RemoveAllPostsFromCurrentUserSavedPosts(c.Request.Context(), requester.GetUserId())
		if err != nil {
			panic(err)
		}
		postCount := 0

		c.JSON(http.StatusOK, common.NewFullResponse("Removed all posts from saved list", posts, &postCount, nil))
	}
}
