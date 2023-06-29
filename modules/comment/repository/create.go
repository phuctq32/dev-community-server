package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo *commentRepository) Create(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error) {
	comment := &entity.Comment{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content:                data.Content,
		PostId:                 data.PostId,
		AuthorId:               data.AuthorId,
		ParentCommentId:        data.ParentCommentId,
		IsApprovedByPostAuthor: false,
	}

	result, err := repo.commentColl.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	*comment.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return comment, nil
}
