package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/comment/business"
	"dev_community_server/modules/comment/entity"
	"dev_community_server/modules/comment/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type CommentBusiness interface {
	CreateComment(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error)
}

type commentHandler struct {
	business CommentBusiness
}

func NewCommentHandler(appCtx appctx.AppContext) *commentHandler {
	cmtRepo := repository.NewCommentRepository(appCtx.GetMongoDBConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetMongoDBConnection())

	biz := business.NewCommentBusiness(cmtRepo, userRepo)

	return &commentHandler{business: biz}
}
