package entity

type UserResetPassword struct {
	Password        *string `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmPassword *string `json:"confirm_password" validate:"eqfield=Password"`
}
