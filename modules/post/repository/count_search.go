package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *postRepository) CountSearch(ctx context.Context, searchTerm *string) (*int, error) {
	filter := bson.M{"$text": bson.M{"$search": searchTerm}}
	cursor, err := repo.postColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var posts []entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, common.NewServerError(err)
	}

	postCount := len(posts)
	return &postCount, nil
}
