package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *commentHandler) GetReplies(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct {
			CommentId string `json:"commentId" validate:"mongodb"`
		}{}
		params.CommentId = c.Param("id")

		if err := appCtx.GetValidator().Validate(&params); err != nil {
			panic(common.NewValidationError(err))
		}

		replies, err := hdl.business.GetReplies(c.Request.Context(), params.CommentId)
		if err != nil {
			panic(err)
		}
		replyCount := len(replies)

		c.JSON(http.StatusOK, common.NewFullResponse("", replies, &replyCount, nil))
	}
}
