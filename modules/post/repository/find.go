package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *postRepository) Find(
	ctx context.Context,
	filter map[string]interface{},
	pagination *common.Pagination,
) ([]entity.Post, error) {
	// Sort by newest
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	// Pagination
	if pagination != nil {
		opts.SetLimit(int64(pagination.Limit)).SetSkip(int64((pagination.Page - 1) * (pagination.Limit)))
	}

	cursor, err := repo.postColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	if posts == nil {
		return []entity.Post{}, nil
	}

	return posts, nil
}
