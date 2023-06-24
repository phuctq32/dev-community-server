package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (repo *postRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Post, error) {
	if err := common.ConvertFieldToObjectId(filter, map[string]string{"id": "_id"}); err != nil {
		return nil, err
	}

	var post entity.Post
	if err := repo.postColl.FindOne(ctx, filter).Decode(&post); err != nil {
		return nil, common.NewServerError(err)
	}

	return &post, nil
}
