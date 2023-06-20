package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/hasher"
	postEntity "dev_community_server/modules/post/entity"
	"dev_community_server/modules/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, data *entity.UserCreate) error
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.User, error)
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type PostRepository interface {
	Find(ctx context.Context, filter common.Filter) ([]*postEntity.Post, error)
}

type userBusiness struct {
	userRepo UserRepository
	postRepo PostRepository
	hash     hasher.MyHash
}

func NewUserBusiness(repo UserRepository, postRepo PostRepository, hash hasher.MyHash) *userBusiness {
	return &userBusiness{userRepo: repo, postRepo: postRepo, hash: hash}
}
