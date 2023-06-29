package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetCurrentUserSavedPosts(ctx context.Context, userId string) ([]entity.Post, error) {
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
	_ = common.AppendInListIdQuery(postFilter, "id", currentUser.SavedPostIds)
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
