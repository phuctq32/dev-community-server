package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (repo *postRepository) Create(ctx context.Context, data *entity.PostCreate) error {
	post := entity.NewPost(data)

	if _, err := repo.postColl.InsertOne(ctx, &post); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
