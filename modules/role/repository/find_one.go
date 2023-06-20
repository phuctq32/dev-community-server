package repository

import (
	"context"
	"dev_community_server/modules/role/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *roleRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Role, error) {
	if id, ok := filter["id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter["_id"] = objId
		delete(filter, "id")
	}

	var role entity.Role
	if err := repo.roleColl.FindOne(ctx, filter).Decode(&role); err != nil {
		return nil, err
	}

	return &role, nil
}
