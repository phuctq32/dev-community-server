package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo *commentRepository) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	// Convert to object id
	postOid, err := common.ToObjectId(comment.PostId)
	if err != nil {
		return nil, err
	}
	authorOid, err := common.ToObjectId(comment.AuthorId)
	if err != nil {
		return nil, err
	}

	insertData := &entity.CommentInsert{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content:                comment.Content,
		PostId:                 postOid,
		AuthorId:               authorOid,
		IsApprovedByPostAuthor: comment.IsApprovedByPostAuthor,
		UpVotes:                comment.UpVotes,
		DownVotes:              comment.DownVotes,
	}

	if comment.ParentCommentId != nil {
		parentCommentOid, err := common.ToObjectId(*comment.ParentCommentId)
		if err != nil {
			return nil, err
		}
		insertData.ParentCommentId = &parentCommentOid
	}

	result, err := repo.commentColl.InsertOne(ctx, &insertData)
	if err != nil {
		return nil, err
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	comment.Id = &insertedId
	comment.CreatedAt = insertData.CreatedAt
	comment.UpdatedAt = insertData.UpdatedAt

	return comment, nil
}
