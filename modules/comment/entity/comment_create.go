package entity

type CommentCreate struct {
	Content         *string `json:"content" validate:"required,min=6"`
	PostId          *string `json:"post_id"`
	ParentCommentId *string `json:"parent_comment_id"`
	AuthorId        *string
}
