package api

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/post/business"
	"dev_community_server/modules/post/entity"
	"dev_community_server/modules/post/repository"
	repository4 "dev_community_server/modules/tag/repository"
	repository3 "dev_community_server/modules/topic/repository"
	repository2 "dev_community_server/modules/user/repository"
)

type PostBusiness interface {
	CreatePost(ctx context.Context, data *entity.PostCreate) (*entity.Post, error)
	GetPosts(ctx context.Context, filter map[string]interface{}, pagination *common.Pagination) ([]entity.Post, *common.PaginationInformation, error)
	UpdatePost(ctx context.Context, data *entity.PostUpdate) (*entity.Post, error)
	GetPostById(ctx context.Context, id *string) (*entity.Post, error)
	SearchPosts(ctx context.Context, searchTerm *string, pagination *common.Pagination) ([]entity.Post, *common.PaginationInformation, error)
}

type postHandler struct {
	business PostBusiness
}

func NewPostHandler(appCtx appctx.AppContext) *postHandler {
	postRepo := repository.NewPostRepository(appCtx.GetMongoDBConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetMongoDBConnection())
	topicRepo := repository3.NewTopicRepository(appCtx.GetMongoDBConnection())
	tagRepo := repository4.NewTagRepository(appCtx.GetMongoDBConnection())
	biz := business.NewPostBusiness(postRepo, userRepo, topicRepo, tagRepo)

	return &postHandler{business: biz}
}
