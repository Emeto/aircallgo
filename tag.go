package aircallgo

type Tag struct {
	ID          int64  `json:"id"`
	DirectLink  string `json:"direct_link"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}
