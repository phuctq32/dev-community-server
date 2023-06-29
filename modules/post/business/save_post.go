package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) SavePost(ctx context.Context, postId string, userId string) ([]entity.Post, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", userId)
	currentUser, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	postFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(postFilter, "id", postId); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}
	if post.Status == entity.Pending {
		return nil, common.NewCustomBadRequestError("Post is pending")
	}

	for _, savedPostId := range currentUser.SavedPostIds {
		if savedPostId == postId {
			return nil, common.NewCustomBadRequestError("Post already exists in saved posts")
		}
	}

	updatedSavedPosts := append(currentUser.SavedPostIds, postId)
	_, err = biz.userRepo.Update(ctx, userFilter, map[string]interface{}{"saved_post_ids": updatedSavedPosts})
	if err != nil {
		return nil, err
	}

	_ = common.AppendInListIdQuery(postFilter, "id", updatedSavedPosts)
	posts, err := biz.postRepo.Find(ctx, postFilter, nil)
	if err != nil {
		return nil, err
	}

	for i := range posts {
		if err = biz.SetComputedDataForPostInList(ctx, &posts[i]); err != nil {
			return nil, err
		}
	}

	return posts, nil
}
