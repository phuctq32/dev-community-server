package repository

import (
	"context"
	"dev_community_server/common"
	"log"
)

func (repo *commentRepository) Count(ctx context.Context, filter map[string]interface{}) (*int, error) {
	postCount, err := repo.commentColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	result := int(postCount)
	log.Println(result)

	return &result, nil
}
