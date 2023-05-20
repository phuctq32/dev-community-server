package entity

type PostCreate struct {
	Title    string   `json:"title" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	Images   []string `json:"images,omitempty"`
	AuthorId string
}
