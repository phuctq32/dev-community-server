package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPostById(ctx context.Context, id *string) (*entity.Post, error) {
	post, err := biz.postRepo.FindOne(ctx, map[string]interface{}{"id": *id})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}

	if err = biz.SetComputedData(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
