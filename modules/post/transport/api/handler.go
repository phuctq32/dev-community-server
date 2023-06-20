package api

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/business"
	"dev_community_server/modules/post/entity"
	"dev_community_server/modules/post/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type PostBusiness interface {
	CreatePost(ctx context.Context, data *entity.PostCreate) error
	GetPosts(ctx context.Context, filter common.Filter) ([]*entity.Post, error)
	UpdatePost(ctx context.Context, data *entity.PostUpdate) error
	GetPostById(ctx context.Context, id *string) (*entity.Post, error)
}

type postHandler struct {
	appCtx   appctx.AppContext
	business PostBusiness
}

func NewPostHandler(appCtx appctx.AppContext) *postHandler {
	postRepo := repository.NewPostRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	biz := business.NewPostBusiness(postRepo, userRepo)

	return &postHandler{appCtx: appCtx, business: biz}
}
