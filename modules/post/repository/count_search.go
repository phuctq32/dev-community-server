package repository

import (
	"context"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *postRepository) CountSearch(ctx context.Context, searchTerm *string) (*int, error) {
	filter := bson.M{"status": entity.Approved, "$text": bson.M{"$search": searchTerm}}
	return repo.Count(ctx, filter)
}
