package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (hdl *postHandler) GetPostById() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := strings.TrimSpace(c.Param("id"))
		post, err := hdl.business.GetPostById(c.Request.Context(), &postId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"post": post,
		})
	}
}
