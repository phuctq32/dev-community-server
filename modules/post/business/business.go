package business

import (
	"context"
	"dev_community_server/common"
	entity4 "dev_community_server/modules/comment/entity"
	"dev_community_server/modules/post/entity"
	entity5 "dev_community_server/modules/role/entity"
	entity2 "dev_community_server/modules/tag/entity"
	entity3 "dev_community_server/modules/topic/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type PostRepository interface {
	Create(ctx context.Context, data *entity.Post) (*entity.Post, error)
	Count(ctx context.Context, filter map[string]interface{}) (*int, error)
	Find(ctx context.Context, filter map[string]interface{}, pagination *common.Pagination) ([]entity.Post, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Post, error)
	Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Post, error)
	Search(ctx context.Context, searchTerm *string, pagination *common.Pagination) ([]entity.Post, error)
	CountSearch(ctx context.Context, searchTerm *string) (*int, error)
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
	Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*userEntity.User, error)
}

type CommentRepository interface {
	Count(ctx context.Context, filter map[string]interface{}) (*int, error)
	Find(ctx context.Context, filter map[string]interface{}) ([]entity4.Comment, error)
}

type TopicRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity3.Topic, error)
	Find(ctx context.Context, filter map[string]interface{}) ([]entity3.Topic, error)
}

type TagRepository interface {
	Create(ctx context.Context, tag *entity2.Tag) (*entity2.Tag, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Tag, error)
}

type RoleRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity5.Role, error)
}

type postBusiness struct {
	postRepo    PostRepository
	userRepo    UserRepository
	commentRepo CommentRepository
	topicRepo   TopicRepository
	tagRepo     TagRepository
	roleRepo    RoleRepository
}

func NewPostBusiness(postRepo PostRepository, userRepo UserRepository, cmtRepo CommentRepository, topicRepo TopicRepository, tagRepo TagRepository, roleRepo RoleRepository) *postBusiness {
	return &postBusiness{postRepo: postRepo, userRepo: userRepo, commentRepo: cmtRepo, topicRepo: topicRepo, tagRepo: tagRepo, roleRepo: roleRepo}
}
