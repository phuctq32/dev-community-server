package middlewares

import (
	"dev_community_server/common"
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckRole(roles []common.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.ReqUser).(common.Requester)
		for _, role := range roles {
			if requester.GetRoleType() != role {
				panic(common.NewNoPermissionError(errors.New("Role access denied")))
			}
		}

		c.Next()
	}
}
