package business

type PostRepository interface {
}

type postBusiness struct {
	postRepo PostRepository
}

func NewPostBusiness(postRepo PostRepository) *postBusiness {
	return &postBusiness{postRepo: postRepo}
}
