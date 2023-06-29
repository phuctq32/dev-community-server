package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *commentHandler) DownVote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct {
			CommentId string `json:"postId" validate:"required,mongodb"`
		}{}
		params.CommentId = c.Param("id")

		if err := appCtx.GetValidator().Validate(&params); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		cmt, err := hdl.business.DownVote(c.Request.Context(), params.CommentId, requester.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("", cmt))
	}
}
