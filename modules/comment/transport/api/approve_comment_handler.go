package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *commentHandler) ApproveComment(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct {
			CommentId string `json:"commentId" validate:"required,mongodb"`
		}{}
		params.CommentId = c.Param("id")

		if err := appCtx.GetValidator().Validate(&params); err != nil {
			panic(common.NewValidationError(err))
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)

		cmt, err := hdl.business.ApproveComment(c.Request.Context(), params.CommentId, requester.GetUserId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleResponse("Approved comment", cmt))
	}
}
