package business

import (
	"context"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) CreateComment(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error) {
	comment, err := biz.commentRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
