package entity

type UserChangePassword struct {
	OldPassword        *string `json:"old_password" validate:"required"`
	NewPassword        *string `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmNewPassword *string `json:"confirm_password" validate:"eqfield=Password"`
}
