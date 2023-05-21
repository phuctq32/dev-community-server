package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	commentColl *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) *commentRepository {
	return &commentRepository{commentColl: db.Collection("comments")}
}
