package repository

import (
	"context"
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *commentRepository) Create(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error) {
	comment := entity.NewComment(data)

	result, err := repo.commentColl.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.Id = result.InsertedID.(primitive.ObjectID)

	return comment, nil
}
