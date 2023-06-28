package repository

import (
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepository struct {
	postColl *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *postRepository {
	return &postRepository{postColl: db.Collection(new(entity.Post).CollectionName())}
}
