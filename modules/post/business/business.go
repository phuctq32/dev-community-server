package business

import (
	"context"
	"dev_community_server/modules/post/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type PostRepository interface {
	Create(ctx context.Context, data *entity.PostCreate) error
	Find(ctx context.Context, filter entity.Filter) ([]*entity.Post, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Post, error)
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
}

type postBusiness struct {
	postRepo PostRepository
	userRepo UserRepository
}

func NewPostBusiness(postRepo PostRepository, userRepo UserRepository) *postBusiness {
	return &postBusiness{postRepo: postRepo, userRepo: userRepo}
}
