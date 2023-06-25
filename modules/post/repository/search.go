package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *postRepository) Search(
	ctx context.Context,
	searchTerm *string,
	pagination *common.Pagination,
) ([]entity.Post, error) {
	opts := options.Find()
	// Pagination
	if pagination != nil {
		opts.SetLimit(int64(*pagination.Limit)).SetSkip(int64((*pagination.Page - 1) * (*pagination.Limit)))
	}

	// Text search
	filter := bson.M{"$text": bson.M{"$search": searchTerm}}

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
