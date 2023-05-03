package entity

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
)

type UserCreate struct {
	Email           string      `json:"email" validate:"required,email"`
	FirstName       string      `json:"first_name" validate:"required"`
	LastName        string      `json:"last_name" validate:"required"`
	Password        string      `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmPassword string      `json:"confirm_password" validate:"eqfield=Password"`
	Birthday        common.Date `json:"birthday"`
}

func (user UserCreate) Validate(appCtx appctx.AppContext) []*common.ValidationError {
	if err := appCtx.GetValidator().Struct(&user); err != nil {
		return common.ValidationErrorsConverter(err)
	}

	return nil
}
