package repository

import (
	"context"
	"dev_community_server/modules/role/entity"
	"dev_community_server/modules/role/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Role, error)
}

type userRepository struct {
	userColl *mongo.Collection
	roleRepo RoleRepository
}

func NewUserRepository(db *mongo.Database) *userRepository {
	roleRepo := repository.NewRoleRepository(db)
	return &userRepository{userColl: db.Collection("users"), roleRepo: roleRepo}
}
