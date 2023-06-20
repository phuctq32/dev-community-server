package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *roleHandler) GetRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := hdl.biz.GetRoles(c.Request.Context())
		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"roles": roles,
		})
	}
}
