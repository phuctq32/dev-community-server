package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPosts(ctx context.Context, filter common.Filter) ([]entity.Post, error) {
	posts, err := biz.postRepo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
