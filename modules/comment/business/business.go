package business

import (
	"context"
	"dev_community_server/modules/comment/entity"
	userEntity "dev_community_server/modules/user/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, data *entity.CommentCreate) (*entity.Comment, error)
}

type UserRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*userEntity.User, error)
}

type commentBusiness struct {
	commentRepo CommentRepository
	userRepo    UserRepository
}

func NewCommentBusiness(commentRepo CommentRepository, userRepo UserRepository) *commentBusiness {
	return &commentBusiness{commentRepo: commentRepo, userRepo: userRepo}
}
