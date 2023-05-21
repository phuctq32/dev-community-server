package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *userHandler) GetPostsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := hdl.business.GetPostsByUserId(c.Request.Context(), c.Param("userId"))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": res,
		})
	}
}
