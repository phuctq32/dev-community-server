package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetCurrentUserSavedPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", (*user).GetUserId())
	currentUser, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, nil, err
	}
	if currentUser == nil {
		return nil, nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	postFilter := map[string]interface{}{}
	_ = common.AppendInListIdQuery(postFilter, "id", currentUser.SavedPostIds)
	posts, err := biz.postRepo.Find(ctx, postFilter, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount := len(posts)

	for i := range posts {
		if err := biz.SetComputedData(ctx, &posts[i]); err != nil {
			return nil, nil, err
		}
	}

	totalPage := totalPostCount / (pagination.Limit)
	if totalPostCount%pagination.Limit > 0 {
		totalPage++
	}
	paginationInfo := &common.PaginationInformation{
		PerPage:   &pagination.Limit,
		Page:      &pagination.Page,
		TotalPage: &totalPage,
	}
	return posts, paginationInfo, nil
}
