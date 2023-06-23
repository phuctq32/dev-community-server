package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/hasher"
	postEntity "dev_community_server/modules/post/entity"
	postRepository "dev_community_server/modules/post/repository"
	"dev_community_server/modules/user/business"
	"dev_community_server/modules/user/entity"
	"dev_community_server/modules/user/repository"
)

type UserBusiness interface {
	GetUserById(ctx context.Context, id string) (user *entity.User, err error)
	UpdateUser(ctx context.Context, id string, data *entity.UserUpdate) error
	ChangePassword(ctx context.Context, id string, user *entity.UserChangePassword) error
	GetPostsByUserId(ctx context.Context, userId string) ([]*postEntity.Post, error)
}

type userHandler struct {
	business UserBusiness
}

func NewUserHandler(appCtx appctx.AppContext) *userHandler {
	userRepo := repository.NewUserRepository(appCtx.GetMongoDBConnection())
	postRepo := postRepository.NewPostRepository(appCtx.GetMongoDBConnection())
	hash := hasher.NewBcryptHash(10)

	biz := business.NewUserBusiness(userRepo, postRepo, hash)
	return &userHandler{business: biz}
}
