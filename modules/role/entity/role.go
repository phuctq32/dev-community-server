package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson"
)

type Role struct {
	common.MongoId `bson:",inline" json:",inline"`
	Name           string          `bson:"name" json:"name"`
	Type           common.RoleType `bson:"type" json:"-"`
}

func (*Role) CollectionName() string { return "roles" }

func (role *Role) MarshalBSON() ([]byte, error) {
	return bson.Marshal(role)
}
