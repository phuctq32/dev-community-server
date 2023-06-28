package repository

import (
	"dev_community_server/modules/role/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type roleRepository struct {
	roleColl *mongo.Collection
}

func NewRoleRepository(db *mongo.Database) *roleRepository {
	return &roleRepository{roleColl: db.Collection(new(entity.Role).CollectionName())}
}
