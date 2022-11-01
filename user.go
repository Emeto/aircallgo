package aircallgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

type UsersAvailabilities struct {
	Meta  Meta               `json:"meta"`
	Users []UserAvailability `json:"users"`
}

type UserAvailability struct {
	ID           int64  `json:"id"`
	Availability string `json:"availability"`
}

type createUserPayload struct {
	email              string
	firstName          string
	lastName           string
	availabilityStatus string
	isAdmin            bool
}

type updateUserPayload struct {
	firstName          string
	lastName           string
	availabilityStatus string
}

func GetUser(UserID int64) *User {
	client := (*Client[User])(newClient())
	user := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10), http.MethodGet, nil)
	return &user.data
}

func GetUsers() *Users {
	client := (*Client[Users])(newClient())
	users := client.MakeRequest("users", http.MethodGet, nil)
	return &users.data
}

func CreateUser(Email string, FirstName string, LastName string, AvailabilityStatus string, IsAdmin bool) *User {
	if len(Email) == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(CreateUser): Email cannot be empty")
		return nil
	}
	if len(FirstName) == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(CreateUser): FirstName cannot be empty")
		return nil
	}
	if len(LastName) == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(CreateUser): LastName cannot be empty")
		return nil
	}
	if !isValidAvailabilityStatus(AvailabilityStatus) {
		AvailabilityStatus = "available"
	}
	client := (*Client[User])(newClient())
	payload, err := json.Marshal(createUserPayload{
		email:              Email,
		firstName:          FirstName,
		lastName:           LastName,
		availabilityStatus: AvailabilityStatus,
		isAdmin:            IsAdmin,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "aircallgo(CreateUser): unable to marshal json payload")
		panic(err)
	}
	r := bytes.NewReader(payload)
	user := client.MakeRequest("users", http.MethodPost, r)
	return &user.data
}

func UpdateUser(UserID int64, FirstName string, LastName string, AvailabilityStatus string) *User {
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(UpdateUser): target user ID cannot be 0")
		return nil
	}
	if !isValidAvailabilityStatus(AvailabilityStatus) {
		AvailabilityStatus = "available"
	}
	client := (*Client[User])(newClient())
	payload, err := json.Marshal(updateUserPayload{
		firstName:          FirstName,
		lastName:           LastName,
		availabilityStatus: AvailabilityStatus,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "aircallgo(UpdateUser): unable to marshal json payload")
		panic(err)
	}
	r := bytes.NewReader(payload)
	user := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10), http.MethodPut, r)
	return &user.data
}

func DeleteUser(UserID int64) bool {
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(DeleteUser): target user ID cannot be 0")
	}
	client := newClient()
	res := client.MakeRequest("users", http.MethodDelete, nil)
	return res.StatusCode == http.StatusNoContent
}

func GetAvailabilities() *UsersAvailabilities {
	client := (*Client[UsersAvailabilities])(newClient())
	ua := client.MakeRequest("users/availabilities", http.MethodGet, nil)
	return &ua.data
}

func GetAvailabilitiesWithFilters(From string, To string, Order string) *UsersAvailabilities {
	es := "users/availabilities?"
	if len(From) > 0 {
		es += "from=" + From
	}
	if len(To) > 0 {
		if len(From) > 0 {
			es += "&"
		}
		es += "to=" + To
	}
	if len(Order) > 0 {
		if len(From) > 0 || len(To) > 0 {
			es += "&"
		}
		es += "order=" + Order
	}
	client := (*Client[UsersAvailabilities])(newClient())
	ua := client.MakeRequest(es, http.MethodGet, nil)
	return &ua.data
}

func GetUserAvailability(UserID int64) *struct {
	Availability string `json:"availability"`
} {
	client := (*Client[struct {
		Availability string `json:"availability"`
	}])(newClient())
	ua := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10)+"/availability", http.MethodGet, nil)
	return &ua.data
}
