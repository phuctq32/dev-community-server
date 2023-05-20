package api

import (
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/business"
	"dev_community_server/modules/post/repository"
)

type PostBusiness interface {
}

type postHandler struct {
	appCtx   appctx.AppContext
	business PostBusiness
}

func NewPostHandler(appCtx appctx.AppContext) *postHandler {
	postRepo := repository.NewPostRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	biz := business.NewPostBusiness(postRepo)

	return &postHandler{appCtx: appCtx, business: biz}
}
