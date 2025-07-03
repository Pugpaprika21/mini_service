package response

type DataService struct {
	TotalRows *int        `json:"total_rows"`
	Data      interface{} `json:"data"`
}
