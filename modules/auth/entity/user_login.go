package entity

import (
	"dev_community_server/common"
)

var (
	ErrorLoginInvalid = common.NewCustomBadRequestError("Email or password invalid")
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
