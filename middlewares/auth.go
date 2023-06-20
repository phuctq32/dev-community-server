package middlewares

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/components/jwt"
	repository2 "dev_community_server/modules/role/repository"
	"dev_community_server/modules/user/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authorize(appCtx appctx.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewJwtProvider(*appCtx.GetSecretKey())

	return func(c *gin.Context) {
		// Extract token from header
		parts := strings.Split(c.GetHeader("Authorization"), " ")
		if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
			panic(common.NewNoPermissionError(errors.New("Wrong authorization header")))
		}

		userRepo := repository.NewUserRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
		roleRepo := repository2.NewRoleRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())

		payload, err := tokenProvider.Decode(parts[1])
		if err != nil {
			panic(err)
		}

		user, err := userRepo.FindOne(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		role, err := roleRepo.FindOne(c.Request.Context(), map[string]interface{}{"id": user.RoleId.Hex()})
		if err != nil {
			panic(err)
		}
		user.RoleType = role.Type

		if !user.IsVerified {
			panic(common.NewCustomBadRequestError("user not verified"))
		}

		c.Set(common.ReqUser, user)
		c.Next()
	}
}
