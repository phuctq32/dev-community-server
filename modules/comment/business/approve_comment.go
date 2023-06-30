package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) ApproveComment(ctx context.Context, cmtId, userId string) (*entity.Comment, error) {
	cmtFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(cmtFilter, "id", cmtId); err != nil {
		return nil, err
	}
	cmt, err := biz.commentRepo.FindOne(ctx, cmtFilter)
	if err != nil {
		return nil, err
	}
	if cmt == nil {
		return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
	}
	if cmt.ParentCommentId != nil {
		return nil, common.NewCustomBadRequestError("Cannot approve a reply")
	}

	postFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(postFilter, "id", cmt.PostId); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}

	userFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(userFilter, "id", userId); err != nil {
		return nil, err
	}
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}
	if *user.Id != post.AuthorId {
		return nil, common.NewCustomBadRequestError("User must be post author to approve comment")
	}

	existingApprovedCmtFilter := map[string]interface{}{"is_approved_by_post_author": true}
	_ = common.AppendIdQuery(existingApprovedCmtFilter, "post_id", *post.Id)
	existingApprovedCmt, err := biz.commentRepo.FindOne(ctx, existingApprovedCmtFilter)
	if err != nil {
		return nil, err
	}
	if existingApprovedCmt != nil {
		if *existingApprovedCmt.Id == *cmt.Id {
			return nil, common.NewCustomBadRequestError("Comment already approved before")
		}
		return nil, common.NewCustomBadRequestError("Can not approve multiple comments in a post")
	}

	updatedCmt, err := biz.commentRepo.Update(ctx, cmtFilter, map[string]interface{}{"is_approved_by_post_author": true})
	if err != nil {
		return nil, err
	}

	if err = biz.SetComputedDataForCommentInList(ctx, updatedCmt); err != nil {
		return nil, err
	}

	return updatedCmt, nil
}
