package entity

type CommentCreate struct {
	Content  *string `json:"content" validate:"required,min=6"`
	AuthorId *string
}
