package business

import (
	"context"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) SearchPost(ctx context.Context, filter entity.Filter) ([]*entity.Post, error) {
	posts, err := biz.postRepo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
