package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) RemovePostFromSavedPosts(ctx context.Context, postId string, userId string) ([]entity.Post, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", userId)
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
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

	isContain := false
	removeIndex := 0
	for i, id := range user.SavedPostIds {
		if postId == id {
			removeIndex = i
			isContain = true
			break
		}
	}
	if !isContain {
		return nil, common.NewCustomBadRequestError("post is not contain in saved posts")
	}

	updatedSavedPosts := append(user.SavedPostIds[:removeIndex], user.SavedPostIds[removeIndex+1:]...)
	_, err = biz.userRepo.Update(ctx, userFilter, map[string]interface{}{"saved_posts_ids": updatedSavedPosts})
	if err != nil {
		return nil, err
	}

	_ = common.AppendInListIdQuery(postFilter, "id", updatedSavedPosts)
	posts, err := biz.postRepo.Find(ctx, postFilter, nil)
	if err != nil {
		return nil, err
	}

	for i := range posts {
		if err = biz.SetComputedData(ctx, &posts[i]); err != nil {
			return nil, err
		}
	}
	return posts, nil
}
