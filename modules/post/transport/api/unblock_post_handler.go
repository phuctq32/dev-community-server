package api

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *postHandler) UnblockPostHandler(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Data struct {
			PostId *string `json:"postId" validate:"required,mongodb"`
		}
		postId := c.Param("id")
		data := Data{PostId: &postId}

		if err := appCtx.GetValidator().Validate(&data); err != nil {
			panic(common.NewValidationError(err))
		}

		user := c.MustGet(common.ReqUser).(common.Requester)

		post, err := hdl.business.UnblockPost(c.Request.Context(), &postId, &user)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSimpleResponse("Unblocked post", post))
	}
}
