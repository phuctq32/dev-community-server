package entity

type CommentUpdate struct {
	Content   string `json:"content" validate:"required,min=4"`
	CommentId string
	UserId    string
}
