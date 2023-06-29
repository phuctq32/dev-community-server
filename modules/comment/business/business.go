package business

import (
	"context"
	"dev_community_server/modules/comment/entity"
	entity2 "dev_community_server/modules/post/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error)
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
}

type PostRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Post, error)
}

type commentBusiness struct {
	commentRepo CommentRepository
	userRepo    UserRepository
	postRepo    PostRepository
}

func NewCommentBusiness(commentRepo CommentRepository, userRepo UserRepository, postRepo PostRepository) *commentBusiness {
	return &commentBusiness{commentRepo: commentRepo, userRepo: userRepo, postRepo: postRepo}
}
