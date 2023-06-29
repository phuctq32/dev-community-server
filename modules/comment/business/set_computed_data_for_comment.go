package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) SetComputedDataForCommentInList(ctx context.Context, cmt *entity.Comment) error {
	// Set author
	if err := biz.SetAuthorData(ctx, cmt); err != nil {
		return err
	}

	// Set reply count
	if err := biz.SetReplyCount(ctx, cmt); err != nil {
		return err
	}

	// Set score
	biz.SetScore(cmt)

	return nil
}

func (biz *commentBusiness) SetComputedData(ctx context.Context, cmt *entity.Comment) error {
	if err := biz.SetComputedDataForCommentInList(ctx, cmt); err != nil {
		return err
	}

	// Set replies
	if err := biz.SetReplies(ctx, cmt); err != nil {
		return err
	}

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
	repliesFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(repliesFilter, "parent_comment_id", *cmt.Id)
	replies, err := biz.commentRepo.Find(ctx, repliesFilter)
	if err != nil {
		return err
	}

	cmt.Replies = &replies
	return nil
}

func (biz *commentBusiness) SetReplyCount(ctx context.Context, cmt *entity.Comment) error {
	repliesFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(repliesFilter, "parent_comment_id", *cmt.Id)
	replyCount, err := biz.commentRepo.Count(ctx, repliesFilter)
	if err != nil {
		return err
	}

	cmt.ReplyCount = replyCount
	return nil
}

func (biz *commentBusiness) SetScore(cmt *entity.Comment) {
	score := len(*cmt.UpVotes) - len(*cmt.DownVotes)
	cmt.Score = &score
}
