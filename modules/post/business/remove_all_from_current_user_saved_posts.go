package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) RemoveAllPostsFromCurrentUserSavedPosts(ctx context.Context, userId string) ([]entity.Post, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", userId)
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	if _, err = biz.userRepo.Update(ctx, userFilter, map[string]interface{}{"saved_post_ids": []string{}}); err != nil {
		return nil, err
	}

	return []entity.Post{}, nil
}
