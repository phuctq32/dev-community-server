package business

import (
	"context"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPostByUserId(ctx context.Context, userId string) ([]*entity.Post, error) {
	posts, err := biz.postRepo.Find(ctx, map[string]interface{}{"author_id": userId})
	if err != nil {
		return nil, err
	}

	return posts, nil
}
