package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
	entity3 "dev_community_server/modules/topic/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type PostRepository interface {
	Create(ctx context.Context, data *entity.PostCreate) (*entity.Post, error)
	Find(ctx context.Context, filter common.Filter) ([]entity.Post, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Post, error)
	Update(ctx context.Context, id string, data map[string]interface{}) error
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
}

type TopicRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity3.Topic, error)
}

type TagRepository interface {
	Create(ctx context.Context, data *entity2.TagCreate) (*entity2.Tag, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Tag, error)
}

type postBusiness struct {
	postRepo  PostRepository
	userRepo  UserRepository
	topicRepo TopicRepository
	tagRepo   TagRepository
}

func NewPostBusiness(postRepo PostRepository, userRepo UserRepository, topicRepo TopicRepository, tagRepo TagRepository) *postBusiness {
	return &postBusiness{postRepo: postRepo, userRepo: userRepo, topicRepo: topicRepo, tagRepo: tagRepo}
}
