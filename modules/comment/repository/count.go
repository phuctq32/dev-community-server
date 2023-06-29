package repository

import (
	"context"
	"dev_community_server/common"
)

func (repo *commentRepository) Count(ctx context.Context, filter map[string]interface{}) (*int, error) {
	postCount, err := repo.commentColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	result := int(postCount)

	return &result, nil
}
