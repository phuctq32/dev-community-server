package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (repo *postRepository) Count(ctx context.Context, filter map[string]interface{}) (*int, error) {
	cursor, err := repo.postColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var posts []entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, common.NewServerError(err)
	}

	postCount := len(posts)
	return &postCount, nil
}
