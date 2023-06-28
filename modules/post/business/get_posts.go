package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPosts(ctx context.Context, filter map[string]interface{}, pagination *common.Pagination) ([]entity.Post, *common.PaginationInformation, error) {
	if pagination.Limit < 0 {
		pagination.Limit = common.DefaultPage
	}
	if pagination.Page < 1 {
		pagination.Page = common.DefaultPage
	}

	// Just get approved post
	filter["status"] = entity.Approved
	posts, err := biz.postRepo.Find(ctx, filter, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount, err := biz.postRepo.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for i := range posts {
		if err = biz.SetComputedData(ctx, &posts[i]); err != nil {
			return nil, nil, err
		}
	}

	var paginationInfo *common.PaginationInformation
	if pagination != nil {
		totalPage := *totalPostCount / (pagination.Limit)
		if *totalPostCount%pagination.Limit > 0 {
			totalPage++
		}
		paginationInfo = &common.PaginationInformation{
			PerPage:   &pagination.Limit,
			Page:      &pagination.Page,
			TotalPage: &totalPage,
		}
	}

	return posts, paginationInfo, nil
}
