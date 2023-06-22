package api

import (
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (hdl *postHandler) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.PostCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := hdl.appCtx.GetValidator().Validate(data); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		requester := c.MustGet(common.ReqUser).(common.Requester)
		data.AuthorId = requester.GetUserId()
		log.Println(data)

		if err := hdl.business.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSimpleResponse("Created post successfully", nil))
	}
}
