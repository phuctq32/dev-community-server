package entity

type CommentCreate struct {
	Content         string  `json:"content" validate:"required,min=6"`
	PostId          string  `json:"post_id" validate:"required,mongodb"`
	ParentCommentId *string `json:"parent_comment_id,omitempty" validate:"omitempty,mongodb"`
	AuthorId        string
}
