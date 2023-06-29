package entity

import (
	"dev_community_server/common"
)

type Role struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string          `bson:"name" json:"name"`
	Type           common.RoleType `bson:"type" json:"-"`
}

func (*Role) CollectionName() string { return "roles" }
