package entity

import (
	"dev_community_server/common"
	"errors"
)

var (
	ErrorLoginInvalid = common.NewBadRequestError("Email or password invalid", errors.New("Email or password invalid"))
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
