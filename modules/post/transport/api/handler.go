package api

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	repository5 "dev_community_server/modules/comment/repository"
	"dev_community_server/modules/post/business"
	"dev_community_server/modules/post/entity"
	"dev_community_server/modules/post/repository"
	repository6 "dev_community_server/modules/role/repository"
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
	ApprovePostById(ctx context.Context, postId *string, user *common.Requester) (*entity.Post, error)
	BlockPost(ctx context.Context, postId *string, user *common.Requester) (*entity.Post, error)
	UnblockPost(ctx context.Context, postId *string, user *common.Requester) (*entity.Post, error)
	GetPendingPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error)
	GetCurrentUserPendingPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error)
	GetCurrentUserApprovedPosts(ctx context.Context, pagination *common.Pagination, user *common.Requester) ([]entity.Post, *common.PaginationInformation, error)
	GetCurrentUserSavedPosts(ctx context.Context, userId string) ([]entity.Post, error)
	SavePost(ctx context.Context, postId string, userId string) ([]entity.Post, error)
	RemoveAllPostsFromCurrentUserSavedPosts(ctx context.Context, userId string) ([]entity.Post, error)
	RemovePostFromSavedPosts(ctx context.Context, postId string, userId string) ([]entity.Post, error)
	UpVote(ctx context.Context, postId string, userId string) (*entity.Post, error)
	DownVote(ctx context.Context, postId string, userId string) (*entity.Post, error)
	GetTrendingPosts(ctx context.Context, quantity int) ([]entity.Post, error)
}

type postHandler struct {
	business PostBusiness
}

func NewPostHandler(appCtx appctx.AppContext) *postHandler {
	postRepo := repository.NewPostRepository(appCtx.GetMongoDBConnection())
	userRepo := repository2.NewUserRepository(appCtx.GetMongoDBConnection())
	cmtRepo := repository5.NewCommentRepository(appCtx.GetMongoDBConnection())
	topicRepo := repository3.NewTopicRepository(appCtx.GetMongoDBConnection())
	tagRepo := repository4.NewTagRepository(appCtx.GetMongoDBConnection())
	roleRepo := repository6.NewRoleRepository(appCtx.GetMongoDBConnection())
	biz := business.NewPostBusiness(postRepo, userRepo, cmtRepo, topicRepo, tagRepo, roleRepo)

	return &postHandler{business: biz}
}
