package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetCurrentUserApprovedPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error) {
	// Get approved posts
	filter := map[string]interface{}{"status": entity.Approved}
	_ = common.AppendIdQuery(filter, "author_id", (*user).GetUserId())

	posts, err := biz.postRepo.Find(ctx, filter, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount, err := biz.postRepo.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for i := range posts {
		if err = biz.SetComputedDataForPostInList(ctx, &posts[i]); err != nil {
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
