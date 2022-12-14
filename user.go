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

type outboundCallPayload struct {
	numberID int64
	to       string
}

type dialPayload struct {
	to string
}

// GetUser retrieves a single user and returns a User
func GetUser(UserID int64) *User {
	client := (*Client[User])(newClient())
	user := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10), http.MethodGet, nil)
	return &user.data
}

// GetUsers retrieves all users and returns Users
func GetUsers() *Users {
	client := (*Client[Users])(newClient())
	users := client.MakeRequest("users", http.MethodGet, nil)
	return &users.data
}

// CreateUser adds a user to the Aircall instance and sends an email invitation
// to passed Email parameter. Email, FirstName and LastName are mandatory and
// cannot be empty strings. AvailabilityStatus should be one of "available",
// "custom" or "unavailable". Returns a User on success.
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

// UpdateUser updates a user. Returns a User on success.
// AvailabilityStatus should be one of "available", "custom", "unavailable".
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

// DeleteUser removes a user from the Aircall instance. Returns true on success, false otherwise.
func DeleteUser(UserID int64) bool {
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(DeleteUser): target user ID cannot be 0")
	}
	client := newClient()
	res := client.MakeRequest("users", http.MethodDelete, nil)
	return res.StatusCode == http.StatusNoContent
}

// GetAvailabilities retrives all users' availabilities. Returns UsersAvailabilities on success.
func GetAvailabilities() *UsersAvailabilities {
	client := (*Client[UsersAvailabilities])(newClient())
	ua := client.MakeRequest("users/availabilities", http.MethodGet, nil)
	return &ua.data
}

// GetAvailabilitiesWithFilters is the same as the GetAvailabilities method
// but with filter parameters. From and To parameters should be in a UNIX
// timestamp format and Order should be either of "asc" or "desc".
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

// GetUserAvailability retrieves a single user availability.
func GetUserAvailability(UserID int64) *struct {
	Availability string `json:"availability"`
} {
	client := (*Client[struct {
		Availability string `json:"availability"`
	}])(newClient())
	ua := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10)+"/availability", http.MethodGet, nil)
	return &ua.data
}

// StartOutboundCall starts an outbound call on the Aircall softphone linked to the user from UserID.
// To must be a phone number in E.164 international format.
// Returns true on success, false otherwise.
func StartOutboundCall(UserID int64, NumberID int64, To string) bool {
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(StartOutboundCall): target user ID cannot be 0")
	}
	if NumberID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(StartOutboundCall): target number ID cannot be 0")
	}
	if len(To) == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(StartOutboundCall): 'to' parameter cannot be an empty string")
	}
	client := newClient()
	payload, err := json.Marshal(outboundCallPayload{
		numberID: NumberID,
		to:       To,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "aircallgo(StartOutboundCall): unable to marshal json payload")
		panic(err)
	}
	r := bytes.NewReader(payload)
	res := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10)+"/calls", http.MethodPost, r)
	return res.StatusCode == http.StatusNoContent
}

// Dial a number directly to the targeted user aircall phone
// To must be a phone number in E.164 international format
func Dial(UserID int64, To string) bool {
	if len(To) == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(Dial): 'to' parameter cannot be an empty string")
	}
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(Dial): target user ID cannot be 0")
	}
	client := newClient()
	payload, err := json.Marshal(dialPayload{
		to: To,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "aircallgo(Dial): unable to marshal json payload")
		panic(err)
	}
	r := bytes.NewReader(payload)
	res := client.MakeRequest("users/"+strconv.FormatInt(UserID, 10)+"/dial", http.MethodPost, r)
	return res.StatusCode == http.StatusNoContent
}
