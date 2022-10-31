package aircallgo

type Participant struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
