package business

import (
	"context"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPostByUserId(ctx context.Context, userId string) ([]*entity.Post, error) {
	filter := entity.Filter{
		Limit: nil,
		Page:  nil,
		Other: map[string]interface{}{"author_id": userId},
	}
	posts, err := biz.postRepo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
