package entity

import (
	"dev_community_server/common"
)

type UserCreate struct {
	Email           string      `json:"email" validate:"required,email"`
	FirstName       string      `json:"first_name" validate:"required"`
	LastName        string      `json:"last_name" validate:"required"`
	Password        string      `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmPassword string      `json:"confirm_password" validate:"eqfield=Password"`
	Birthday        common.Date `json:"birthday"`
	RoleId          string
	VerifiedToken   string
}
