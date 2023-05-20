package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) UpdatePost(ctx context.Context, data *entity.PostUpdate) error {
	_, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": *data.AuthorId})
	if err != nil {
		return err
	}

	post, err := biz.postRepo.FindOne(ctx, map[string]interface{}{"id": *data.Id})
	if err != nil {
		if err == common.ErrNotFound {
			return common.NewNotFoundError("Post", err)
		}
		return err
	}

	// check if user is author
	if post.AuthorId.Hex() != *data.AuthorId {
		return common.NewCustomBadRequestError("User is not author")
	}

	update, _ := common.StructToMap(data)

	if err = biz.postRepo.Update(ctx, *data.Id, update); err != nil {
		return err
	}

	return nil
}
