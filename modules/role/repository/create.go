package repository

import (
	"context"
	"dev_community_server/modules/role/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *roleRepository) Create(ctx context.Context, data *entity.RoleCreate) (*entity.Role, error) {
	role := entity.NewRole(data)

	result, err := repo.roleColl.InsertOne(ctx, role)
	if err != nil {
		return nil, err
	}
	role.Id = result.InsertedID.(primitive.ObjectID)

	return role, nil
}
