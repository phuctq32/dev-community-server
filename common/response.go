package common

type response struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Count     *int        `json:"count,omitempty"`
	Page      *int        `json:"page,omitempty"`
	PerPage   *int        `json:"per_page,omitempty"`
	TotalPage *int        `json:"total_page,omitempty"`
}

func NewSimpleResponse(message string, data interface{}) *response {
	return &response{Message: message, Data: data, Count: nil, PerPage: nil, Page: nil, TotalPage: nil}
}

func NewFullResponse(message string, data interface{}, count, page, perPage *int) *response {
	return &response{Message: message, Data: data, Count: count, Page: page, PerPage: perPage}
}
