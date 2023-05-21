package business

type CommentRepository interface {
}

type UserRepository interface {
}

type commentBusiness struct {
	commentRepo CommentRepository
	userRepo    UserRepository
}

func NewCommentBusiness(commentRepo CommentRepository, userRepo UserRepository) *commentBusiness {
	return &commentBusiness{commentRepo: commentRepo, userRepo: userRepo}
}
