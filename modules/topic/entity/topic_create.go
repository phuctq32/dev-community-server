package entity

type TopicCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
