package repository

import (
	"context"
)

func (repo *postRepository) Count(ctx context.Context, filter map[string]interface{}) (*int, error) {
	postCount, err := repo.postColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	result := int(postCount)

	return &result, nil
}
