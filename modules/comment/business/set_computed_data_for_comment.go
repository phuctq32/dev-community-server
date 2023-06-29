package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) SetComputedData(ctx context.Context, cmt *entity.Comment) error {
	// Set author
	if err := biz.SetAuthorData(ctx, cmt); err != nil {
		return err
	}

	// Set replies
	if err := biz.SetReplies(ctx, cmt); err != nil {
		return err
	}

	// Set score
	biz.SetScore(cmt)

	return nil
}

func (biz *commentBusiness) SetAuthorData(ctx context.Context, cmt *entity.Comment) error {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", cmt.AuthorId)
	author, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return err
	}
	if author == nil {
		return common.NewNotFoundError("User", common.ErrNotFound)
	}
	cmt.Author = author
	return nil
}

func (biz *commentBusiness) SetReplies(ctx context.Context, cmt *entity.Comment) error {
	return nil
}

func (biz *commentBusiness) SetScore(cmt *entity.Comment) {
	cmt.Score = len(cmt.UpVotes) - len(cmt.DownVotes)
}
