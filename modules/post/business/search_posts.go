package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"strings"
)

func (biz *postBusiness) SearchPosts(ctx context.Context, searchTerm *string, pagination *common.Pagination) ([]entity.Post, *common.PaginationInformation, error) {
	*searchTerm = strings.Replace(*searchTerm, "+", " ", -1)
	*searchTerm = common.RemoveVietnameseAccent(*searchTerm)

	if pagination != nil {
		if *pagination.Limit < 0 {
			*pagination.Limit = common.DefaultPage
		}
		if *pagination.Page < 1 {
			*pagination.Page = common.DefaultPage
		}
	}

	posts, err := biz.postRepo.Search(ctx, searchTerm, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount, err := biz.postRepo.CountSearch(ctx, searchTerm)
	if err != nil {
		return nil, nil, err
	}

	for i := range posts {
		if err := biz.SetComputedData(ctx, &posts[i]); err != nil {
			return nil, nil, err
		}
	}

	var paginationInfo *common.PaginationInformation
	if pagination != nil {
		totalPage := *totalPostCount / (*pagination.Limit)
		if *totalPostCount%*pagination.Limit > 0 {
			totalPage++
		}
		paginationInfo = &common.PaginationInformation{
			PerPage:   pagination.Limit,
			Page:      pagination.Page,
			TotalPage: &totalPage,
		}
	}

	return posts, paginationInfo, nil
}
