package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) SavePost(ctx context.Context, postId *string, user *common.Requester) (*entity.Post, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", (*user).GetUserId())
	currentUser, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	postFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(userFilter, "id", (*user).GetUserId()); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}

	for _, savedPostId := range currentUser.SavedPostIds {
		if savedPostId == *postId {
			return nil, common.NewCustomBadRequestError("Post already exists in saved posts")
		}
	}

	newSavedPostIds := append(currentUser.SavedPostIds, *postId)
	_, err = biz.userRepo.Update(ctx, userFilter, map[string]interface{}{"saved_posts_ids": newSavedPostIds})
	if err != nil {
		return nil, err
	}

	return post, nil
}
