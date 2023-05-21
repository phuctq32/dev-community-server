package api

import (
	"dev_community_server/components/appctx"
	"dev_community_server/modules/comment/business"
	"dev_community_server/modules/comment/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type CommentBusiness interface {
}

type commentHandler struct {
	business CommentBusiness
}

func NewCommentHandler(appCtx appctx.AppContext) *commentHandler {
	cmtRepo := repository.NewCommentRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())

	biz := business.NewCommentBusiness(cmtRepo, userRepo)

	return &commentHandler{business: biz}
}
