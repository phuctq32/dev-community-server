package entity

type Filter struct {
	Limit  *int
	Page   *int
	Search *string
	Other  map[string]interface{}
}
