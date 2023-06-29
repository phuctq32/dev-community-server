package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	entity2 "dev_community_server/modules/post/entity"
)

func (biz *commentBusiness) CreateComment(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", data.AuthorId)
	user, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	postFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(postFilter, "id", data.PostId); err != nil {
		return nil, err
	}
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

	comment := &entity.Comment{
		Content:  data.Content,
		PostId:   *post.Id,
		AuthorId: *user.Id,
	}

	if data.ParentCommentId != nil {
		parentCmtFilter := map[string]interface{}{}
		if err = common.AppendIdQuery(parentCmtFilter, "id", *data.ParentCommentId); err != nil {
			return nil, err
		}
		parentCmt, err := biz.commentRepo.FindOne(ctx, parentCmtFilter)
		if err != nil {
			return nil, err
		}
		if parentCmt == nil {
			return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
		}
		if parentCmt.ParentCommentId != nil {
			return nil, common.NewCustomBadRequestError("Cannot create reply for a reply")
		}
		comment.ParentCommentId = parentCmt.Id
	} else {
		isApproved := false
		upVotes := []string{}
		downVote := []string{}
		comment.IsApprovedByPostAuthor = &isApproved
		comment.UpVotes = &upVotes
		comment.DownVotes = &downVote
	}

	comment, err = biz.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	if data.ParentCommentId == nil {
		if err = biz.SetComputedData(ctx, comment); err != nil {
			return nil, err
		}
	} else {
		if err = biz.SetAuthorData(ctx, comment); err != nil {
			return nil, err
		}
	}

	return comment, nil
}
