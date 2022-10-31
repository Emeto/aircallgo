package aircallgo

import "time"

type Webhook struct {
	ID         int64     `json:"id"`
	DirectLink string    `json:"direct_link"`
	CreatedAt  time.Time `json:"created_at"`
	CustomName string    `json:"custom_name"`
	URL        string    `json:"url"`
	Active     bool      `json:"active"`
	Token      string    `json:"token"`
	Events     []string  `json:"events"`
}
