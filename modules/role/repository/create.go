package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/role/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *roleRepository) Create(ctx context.Context, data *entity.RoleCreate) (*entity.Role, error) {
	role := &entity.Role{
		Name: data.Name,
	}

	result, err := repo.roleColl.InsertOne(ctx, role)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	role.Id = &insertedId

	return role, nil
}
