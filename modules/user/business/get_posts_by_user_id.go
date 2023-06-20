package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *userBusiness) GetPostsByUserId(ctx context.Context, userId string) ([]*entity.Post, error) {
	user, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	filter := common.Filter{
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
