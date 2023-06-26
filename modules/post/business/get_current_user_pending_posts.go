package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetCurrentUserPendingPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error) {
	if (*user).GetRoleType() == common.Administrator || (*user).GetRoleType() == common.Moderator {
		totalPostCount := 0
		return []entity.Post{}, &common.PaginationInformation{PerPage: &pagination.Limit, Page: &pagination.Page, TotalPage: &totalPostCount}, nil
	}

	// Get pending posts have author is requester (just Member, because Admin or Mod's posts automatically approved)
	filter := map[string]interface{}{
		"author_id": (*user).GetUserId(),
		"status":    entity.Pending,
	}

	posts, err := biz.postRepo.Find(ctx, filter, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount, err := biz.postRepo.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for i := range posts {
		if err := biz.SetComputedData(ctx, &posts[i]); err != nil {
			return nil, nil, err
		}
	}

	totalPage := *totalPostCount / (pagination.Limit)
	if *totalPostCount%pagination.Limit > 0 {
		totalPage++
	}
	paginationInfo := &common.PaginationInformation{
		PerPage:   &pagination.Limit,
		Page:      &pagination.Page,
		TotalPage: &totalPage,
	}
	return posts, paginationInfo, nil
}
