package aircallgo

type Integration struct {
	Name         string  `json:"name"`
	CustomName   string  `json:"custom_name"`
	Logo         string  `json:"logo"`
	CompanyID    int64   `json:"company_id"`
	Status       string  `json:"status"`
	Active       bool    `json:"active"`
	NumberIDs    []int64 `json:"number_ids"`
	NumbersCount int64   `json:"numbers_count"`
	User         User    `json:"user"`
}
