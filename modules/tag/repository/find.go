package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
)

func (repo *tagRepository) Find(ctx context.Context, filter map[string]interface{}) ([]entity.Tag, error) {
	cursor, err := repo.tagColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	tags := make([]entity.Tag, 0)
	if err := cursor.All(ctx, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
