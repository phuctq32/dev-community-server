package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	entity2 "dev_community_server/modules/post/entity"
)

func (biz *commentBusiness) UpdateComment(ctx context.Context, data *entity.CommentUpdate) (*entity.Comment, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", data.UserId)
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	cmtFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(cmtFilter, "id", data.CommentId)
	cmt, err := biz.commentRepo.FindOne(ctx, cmtFilter)
	if err != nil {
		return nil, err
	}
	if cmt == nil {
		return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
	}

	if *user.Id != cmt.AuthorId {
		return nil, common.NewCustomBadRequestError("User is not comment author")
	}

	postFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(postFilter, "id", cmt.PostId)
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}
	if post.Status == entity2.Pending {
		return nil, common.NewCustomBadRequestError("Post is in pending status")
	}
	if post.IsBlocked {
		return nil, common.NewCustomBadRequestError("Post is block. Not allow to comment")
	}

	updateData := map[string]interface{}{"content": data.Content}
	updatedCmt, err := biz.commentRepo.Update(ctx, cmtFilter, updateData)
	if err != nil {
		return nil, err
	}

	if err = biz.SetComputedData(ctx, updatedCmt); err != nil {
		return nil, err
	}

	return updatedCmt, nil
}
