package api

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {
	dd, err := time.Parse(`"2006-01-02T15:04:05.000+0000"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(dd.Format("01/02/2006"))

	return nil
}

func (hdl *authHandler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userEntity.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		if err := data.Validate(hdl.appCtx); err != nil {
			panic(common.NewValidationError(err))
			return
		}

		if err := hdl.business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Register successfully!",
		})
	}
}
