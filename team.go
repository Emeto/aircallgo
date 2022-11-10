package aircallgo

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Team struct {
	ID         int64     `json:"id"`
	DirectLink string    `json:"direct_link"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	Users      []User    `json:"users"`
}

type Teams struct {
	Meta  Meta   `json:"meta"`
	Teams []Team `json:"teams"`
}

func GetTeams() *Teams {
	client := (*Client[Teams])(newClient())
	teams := client.MakeRequest("teams", http.MethodGet, nil)
	return &teams.data
}

func GetTeam(TeamID int64) *Team {
	if TeamID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(GetTeam): target team ID cannot be 0")
		return nil
	}
	client := (*Client[Team])(newClient())
	team := client.MakeRequest("teams/"+strconv.FormatInt(TeamID, 10), http.MethodGet, nil)
	return &team.data
}

func DeleteTeam(TeamID int64) bool {
	if TeamID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(DeleteTeam): target team ID cannot be 0")
		return false
	}
	client := newClient()
	response := client.MakeRequest("teams/"+strconv.FormatInt(TeamID, 10), http.MethodDelete, nil)
	return response.StatusCode == http.StatusOK
}

func AddUserToTeam(TeamID int64, UserID int64) *Team {
	if TeamID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(AddUserToTeam): target team ID cannot be 0")
		return nil
	}
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(AddUserToTeam): target user ID cannot be 0")
		return nil
	}
	client := (*Client[Team])(newClient())
	response := client.MakeRequest("teams/"+strconv.FormatInt(TeamID, 10)+"/users/"+strconv.FormatInt(UserID, 10), http.MethodPost, nil)
	return &response.data
}

func DeleteUserFromTeam(TeamID int64, UserID int64) *Team {
	if TeamID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(DeleteUserFromTeam): target team ID cannot be 0")
		return nil
	}
	if UserID == 0 {
		fmt.Fprint(os.Stderr, "aircallgo(DeleteUserFromTeam): target user ID cannot be 0")
		return nil
	}
	client := (*Client[Team])(newClient())
	response := client.MakeRequest("teams/"+strconv.FormatInt(TeamID, 10)+"/users/"+strconv.FormatInt(UserID, 10), http.MethodDelete, nil)
	return &response.data
}
