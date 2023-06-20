package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/role/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *roleRepository) Find(ctx context.Context) ([]*entity.Role, error) {
	cursor, err := repo.roleColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var roles []*entity.Role
	if err := cursor.All(ctx, &roles); err != nil {
		return nil, common.NewServerError(err)
	}

	return roles, nil
}
