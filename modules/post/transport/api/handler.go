package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/business"
	"dev_community_server/modules/post/entity"
	"dev_community_server/modules/post/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type PostBusiness interface {
	CreatePost(ctx context.Context, data *entity.PostCreate) error
	GetPostByUserId(ctx context.Context, userId string) ([]*entity.Post, error)
	GetPosts(ctx context.Context, filter entity.Filter) ([]*entity.Post, error)
	UpdatePost(ctx context.Context, data *entity.PostUpdate) error
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
