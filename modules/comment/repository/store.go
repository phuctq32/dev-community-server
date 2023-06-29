package repository

import (
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	commentColl *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) *commentRepository {
	return &commentRepository{commentColl: db.Collection(new(entity.Comment).CollectionName())}
}
