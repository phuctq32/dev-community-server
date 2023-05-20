package business

import (
	"context"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) CreatePost(ctx context.Context, data *entity.PostCreate) error {
	_, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": data.AuthorId})
	if err != nil {
		return err
	}

	if err = biz.postRepo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
