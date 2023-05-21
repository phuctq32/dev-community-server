package entity

import (
	"dev_community_server/common"
)

type UserUpdate struct {
	Email      *string      `json:"email,omitempty" validate:"omitempty,email" map:"email"`
	FirstName  *string      `json:"first_name,omitempty" map:"first_name"`
	LastName   *string      `json:"last_name,omitempty" map:"last_name"`
	Birthday   *common.Date `json:"birthday,omitempty" map:"-"`
	Avatar     *string      `json:"avatar" map:"avatar"`
	IsVerified *bool        `json:"is_verified,omitempty" map:"is_verified"`
}
