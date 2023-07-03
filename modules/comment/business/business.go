package business

import (
	"context"
	"dev_community_server/modules/comment/entity"
	entity2 "dev_community_server/modules/post/entity"
	entity3 "dev_community_server/modules/role/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, data *entity.Comment) (*entity.Comment, error)
	Find(ctx context.Context, filter map[string]interface{}) ([]entity.Comment, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Comment, error)
	Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (*entity.Comment, error)
	Count(ctx context.Context, filter map[string]interface{}) (*int, error)
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
}

type PostRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Post, error)
}

type RoleRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity3.Role, error)
}

type commentBusiness struct {
	commentRepo CommentRepository
	userRepo    UserRepository
	postRepo    PostRepository
	roleRepo    RoleRepository
}

func NewCommentBusiness(commentRepo CommentRepository, userRepo UserRepository, postRepo PostRepository, roleRepo RoleRepository) *commentBusiness {
	return &commentBusiness{commentRepo: commentRepo, userRepo: userRepo, postRepo: postRepo, roleRepo: roleRepo}
}
