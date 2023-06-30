package business

import (
	"context"
	"dev_community_server/modules/post/entity"
	"sort"
	"time"
)

func (biz *postBusiness) GetTrendingPosts(ctx context.Context, quantity int) ([]entity.Post, error) {
	if quantity < 3 {
		quantity = 5
	}

	// Get the posts which have most viewed in 3 days recent
	allPosts, err := biz.postRepo.Find(ctx, map[string]interface{}{"status": entity.Approved}, nil)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	threeDaysAgo := now.Add(-time.Hour * 24 * 3)
	for i := range allPosts {
		for _, t := range allPosts[i].TimestampViews {
			if t.After(threeDaysAgo) && t.Before(now) {
				allPosts[i].ViewsPerThreeDays++
			}
		}
	}

	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].ViewsPerThreeDays > allPosts[j].ViewsPerThreeDays
	})

	trendingPosts := allPosts
	if len(allPosts) > quantity {
		trendingPosts = allPosts[:quantity]
	}

	for i := range trendingPosts {
		if err = biz.SetComputedDataForPostInList(ctx, &trendingPosts[i]); err != nil {
			return nil, err
		}
	}

	return trendingPosts, err
}
