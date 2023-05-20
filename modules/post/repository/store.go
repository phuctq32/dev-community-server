package repository

import "go.mongodb.org/mongo-driver/mongo"

type postRepository struct {
	postColl *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *postRepository {
	return &postRepository{postColl: db.Collection("posts")}
}
