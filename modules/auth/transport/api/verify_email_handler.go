package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hdl *authHandler) VerifyEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiedToken := c.Param("verifiedToken")

		if err := hdl.business.VerifyEmail(c.Request.Context(), verifiedToken); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Your email is verified!",
		})
	}
}
