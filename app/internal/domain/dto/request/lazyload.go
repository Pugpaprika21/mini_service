package request

type Lazyload struct {
	PageNo      *int    `json:"page_no"`
	PageSize    *int    `json:"page_size"`
	SearchField *string `json:"search_field"`
}
