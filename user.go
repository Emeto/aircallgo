package aircallgo

import (
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID                 int64     `json:"id"`
	DirectLink         string    `json:"direct_link"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	CreatedAt          time.Time `json:"created_at"`
	Available          bool      `json:"available"`
	AvailabilityStatus string    `json:"availability_status"`
	Numbers            []Number  `json:"numbers"`
	Timezone           string    `json:"time_zone"`
	Language           string    `json:"language"`
	WrapUpTime         int64     `json:"wrap_up_time"`
}

type Users struct {
	Meta  Meta   `json:"meta"`
	Users []User `json:"users"`
}

func GetUser(UserID int64) *User {
	client := (*Client[User])(newClient())
	user := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10), http.MethodGet, nil)
	return &user
}

func GetUsers() *Users {
	client := (*Client[Users])(newClient())
	users := client.MakeRequest("users", http.MethodGet, nil)
	return &users
}
