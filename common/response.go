package common

type response struct {
	Message               string      `json:"message"`
	Data                  interface{} `json:"data"`
	Count                 *int        `json:"count,omitempty"`
	PaginationInformation `json:",inline"`
	//Page      *int        `json:"page,omitempty"`
	//PerPage   *int        `json:"per_page,omitempty"`
	//TotalPage *int        `json:"total_page,omitempty"`
}

func NewSimpleResponse(message string, data interface{}) *response {
	return &response{
		Message: message,
		Data:    data,
		Count:   nil,
		PaginationInformation: PaginationInformation{
			PerPage:   nil,
			Page:      nil,
			TotalPage: nil,
		},
	}
}

func NewFullResponse(message string, data interface{}, count *int, paginationInfo *PaginationInformation) *response {
	if paginationInfo == nil {
		paginationInfo = &PaginationInformation{
			PerPage:   nil,
			Page:      nil,
			TotalPage: nil,
		}
	}
	return &response{
		Message:               message,
		Data:                  data,
		Count:                 count,
		PaginationInformation: *paginationInfo,
	}
}
