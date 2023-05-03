package middlewares

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				appErr, ok := err.(*common.AppError)
				if !ok {
					appErr = common.NewServerError(err.(error))
				}

				appErr.Logging()
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(appErr)
				return
			}
		}()

		c.Next()
	}
}
