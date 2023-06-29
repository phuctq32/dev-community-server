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
		return nil, common.NewCustomBadRequestError("Post is pending")
	}

	comment, err := biz.commentRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	if err = biz.SetComputedData(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}
