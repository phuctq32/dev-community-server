package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) GetReplies(ctx context.Context, parentCmtId string) ([]entity.Comment, error) {
	cmtFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(cmtFilter, "id", parentCmtId); err != nil {
		return nil, err
	}
	cmt, err := biz.commentRepo.FindOne(ctx, cmtFilter)
	if err != nil {
		return nil, err
	}
	if cmt == nil {
		return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
	}

	repliesFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(repliesFilter, "parent_comment_id", *cmt.Id)
	replies, err := biz.commentRepo.Find(ctx, repliesFilter)
	if err != nil {
		return nil, err
	}

	for i := range replies {
		if err = biz.SetAuthorData(ctx, &replies[i]); err != nil {
			return nil, err
		}
	}

	return replies, nil
}
