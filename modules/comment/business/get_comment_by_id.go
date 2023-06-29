package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) GetCommentById(ctx context.Context, id string) (*entity.Comment, error) {
	filter := map[string]interface{}{}
	_ = common.AppendIdQuery(filter, "id", id)
	cmt, err := biz.commentRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	if cmt == nil {
		return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
	}

	if err = biz.SetComputedData(ctx, cmt); err != nil {
		return nil, err
	}

	return cmt, err
}
