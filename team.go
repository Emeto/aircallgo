package aircallgo

import (
	"net/http"
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
	client := (*Client[Team])(newClient())
	team := client.MakeRequest("teams/"+strconv.FormatInt(TeamID, 10), http.MethodGet, nil)
	return &team.data
}
