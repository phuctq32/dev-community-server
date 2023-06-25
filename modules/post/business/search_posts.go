package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
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
		// Get author
		author, _ := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": posts[i].AuthorId.Hex()})
		posts[i].Author = author

		// Get topic
		topic, _ := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": posts[i].TopicId.Hex()})
		posts[i].Topic = topic

		// Get tags
		posts[i].Tags = make([]entity2.Tag, len(posts[i].TagIds))
		for j, id := range posts[i].TagIds {
			tag, _ := biz.tagRepo.FindOne(ctx, map[string]interface{}{"id": id.Hex()})
			posts[i].Tags[j] = *tag
		}

		// Get comments (not include replies) and count total comments (included replies)

		// Calc score
		posts[i].Score = len(posts[i].UpVotes) - len(posts[i].DownVotes)
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
