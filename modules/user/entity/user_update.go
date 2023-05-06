package entity

import (
	"dev_community_server/common"
)

type UserUpdate struct {
	Email      *string     `json:"email,omitempty" validate:"omitempty,email" map:"email"`
	FirstName  *string     `json:"first_name,omitempty" map:"first_name"`
	LastName   *string     `json:"last_name,omitempty" map:"-"`
	Birthday   common.Date `json:"birthday,omitempty" map:"birthday"`
	IsVerified *bool       `json:"is_verified,omitempty" map:"is_verified"`
}
