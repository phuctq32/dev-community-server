package entity

import (
	"dev_community_server/common"
	"time"
)

type Role struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Name               string          `bson:"name" json:"name"`
	Type               common.RoleType `bson:"type" json:"-"`
}

func NewRole(role *RoleCreate) *Role {
	now := time.Now()
	return &Role{
		ModelCommon: common.ModelCommon{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Name: role.Name,
	}
}
