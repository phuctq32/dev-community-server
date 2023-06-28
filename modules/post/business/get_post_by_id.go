package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPostById(ctx context.Context, id *string) (*entity.Post, error) {
	filter := map[string]interface{}{}
	if err := common.AppendIdQuery(filter, "id", *id); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, filter)
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
