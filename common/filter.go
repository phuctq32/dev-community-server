package common

type Filter struct {
	Limit  *int
	Page   *int
	Search *string
	Other  map[string]interface{}
}

type Pagination struct {
	Limit int
	Page  int
}

type PaginationInformation struct {
	PerPage   *int `json:"per_page,omitempty"`
	Page      *int `json:"page,omitempty"`
	TotalPage *int `json:"total_page,omitempty"`
}

const (
	DefaultLimit int = 10
	DefaultPage  int = 1
)

var DefaultPagination *Pagination = &Pagination{Limit: DefaultLimit, Page: DefaultPage}
