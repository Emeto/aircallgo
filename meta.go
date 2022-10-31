package aircallgo

type Meta struct {
	Count            int64  `json:"count"`
	Total            int64  `json:"total"`
	CurrentPage      int64  `json:"current_page"`
	PerPage          int64  `json:"per_page"`
	NextPageLink     string `json:"next_page_link"`
	PreviousPageLink string `json:"previous_page_link"`
}
