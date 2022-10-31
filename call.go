package aircallgo

type Call struct {
	ID               int64         `json:"id"`
	DirectLink       string        `json:"direct_link"`
	StartedAt        int64         `json:"started_at"`
	AnsweredAt       int64         `json:"answered_at"`
	EndedAt          int64         `json:"ended_at"`
	Duration         int64         `json:"duration"`
	Status           string        `json:"status"`
	Direction        string        `json:"direction"`
	RawDigits        string        `json:"raw_digits"`
	Asset            string        `json:"asset"`
	Recording        string        `json:"recording"`
	Voicemail        string        `json:"voicemail"`
	Archived         bool          `json:"archived"`
	MissedCallReason string        `json:"missed_call_reason"`
	Cost             string        `json:"cost"`
	Number           Number        `json:"number"`
	User             User          `json:"user"`
	Contact          Contact       `json:"contact"`
	AssignedTo       User          `json:"assigned_to"`
	Teams            []Team        `json:"teams"`
	TransferredBy    User          `json:"transferred_by"`
	TransferredTo    User          `json:"transferred_to"`
	Comments         []string      `json:"comments"`
	Tags             []Tag         `json:"tags"`
	Participants     []interface{} `json:"participants"`
}
