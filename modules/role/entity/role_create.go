package entity

type RoleCreate struct {
	Name string `json:"name" validate:"required"`
}
