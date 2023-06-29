package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *commentRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Comment, error) {
	var cmt entity.Comment
	if err := repo.commentColl.FindOne(ctx, filter).Decode(&cmt); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.NewServerError(err)
	}

	return &cmt, nil
}
