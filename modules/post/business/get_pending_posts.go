package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) GetPendingPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error) {
	// Get pending posts
	postFilter := map[string]interface{}{"status": entity.Pending}

	// Admin: All pending posts
	// Moderator: Pending posts of topic which he manages
	if (*user).GetRoleType() == common.Moderator {
		managedTopics, err := biz.topicRepo.Find(ctx, map[string]interface{}{"moderator_ids": (*user).GetUserId()})
		if err != nil {
			return nil, nil, err
		}
		if len(managedTopics) > 0 {
			topicIds := make([]string, len(managedTopics))
			for _, topic := range managedTopics {
				topicIds = append(topicIds, *topic.Id)
			}
			_ = common.AppendInListIdQuery(postFilter, "topic_id", topicIds)
		}
	}

	posts, err := biz.postRepo.Find(ctx, postFilter, pagination)
	if err != nil {
		return nil, nil, err
	}
	totalPostCount, err := biz.postRepo.Count(ctx, postFilter)
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
