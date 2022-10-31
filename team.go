package aircallgo

import "time"

type Team struct {
	ID         int64     `json:"id"`
	DirectLink string    `json:"direct_link"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	Users      []User    `json:"users"`
}
