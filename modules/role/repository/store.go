package repository

import "go.mongodb.org/mongo-driver/mongo"

type roleRepository struct {
	roleColl *mongo.Collection
}

func NewRoleRepository(db *mongo.Database) *roleRepository {
	return &roleRepository{roleColl: db.Collection("roles")}
}
