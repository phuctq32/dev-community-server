package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/comment/business"
	"dev_community_server/modules/comment/entity"
	"dev_community_server/modules/comment/repository"
	repository3 "dev_community_server/modules/post/repository"
	repository4 "dev_community_server/modules/role/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type CommentBusiness interface {
	CreateComment(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error)
	GetCommentById(ctx context.Context, id string) (*entity.Comment, error)
	UpdateComment(ctx context.Context, data *entity.CommentUpdate) (*entity.Comment, error)
	GetReplies(ctx context.Context, parentCmtId string) ([]entity.Comment, error)
	UpVote(ctx context.Context, cmtId string, userId string) (*entity.Comment, error)
	DownVote(ctx context.Context, cmtId string, userId string) (*entity.Comment, error)
	ApproveComment(ctx context.Context, cmtId, userId string) (*entity.Comment, error)
	UnApproveComment(ctx context.Context, cmtId, userId string) (*entity.Comment, error)
}

type commentHandler struct {
	business CommentBusiness
}

func NewCommentHandler(appCtx appctx.AppContext) *commentHandler {
	cmtRepo := repository.NewCommentRepository(appCtx.GetMongoDBConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetMongoDBConnection())
	postRepo := repository3.NewPostRepository(appCtx.GetMongoDBConnection())
	roleRepo := repository4.NewRoleRepository(appCtx.GetMongoDBConnection())

	biz := business.NewCommentBusiness(cmtRepo, userRepo, postRepo, roleRepo)

	return &commentHandler{business: biz}
}
