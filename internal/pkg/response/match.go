package response

type Match struct {
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Status string  `json:"status"`
	Result *string `json:"result,omitempty"`
}
