package entity

type PostUpdate struct {
	Id       *string `map:"-"`
	AuthorId *string `map:"-"`
	Title    *string `json:"title,omitempty" validate:"omitempty,min=4" map:"title"`
	Content  *string `json:"content,omitempty" validate:"omitempty,min=4" map:"content"`
	Images   *string `json:"images,omitempty" map:"images"`
}
