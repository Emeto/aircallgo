package aircallgo

import "time"

type Number struct {
	ID                     int64     `json:"id"`
	DirectLink             string    `json:"direct_link"`
	Name                   string    `json:"name"`
	Digits                 string    `json:"digits"`
	CreatedAt              time.Time `json:"created_at"`
	Country                string    `json:"country"`
	Timezone               string    `json:"time_zone"`
	Open                   bool      `json:"open"`
	AvailabilityStatus     string    `json:"availability_status"`
	IsIVR                  bool      `json:"is_ivr"`
	LiveRecordingActivated bool      `json:"live_recording_activated"`
	Users                  []User    `json:"users"`
	Priority               int64     `json:"priority"`
	Messages               Messages  `json:"messages"`
}

type Messages struct {
	Welcome        string `json:"welcome"`
	Waiting        string `json:"waiting"`
	RingingTone    string `json:"ringing_tone"`
	UnansweredCall string `json:"unanswered_call"`
	AfterHours     string `json:"after_hours"`
	IVR            string `json:"ivr"`
	Voicemail      string `json:"voicemail"`
	Closed         string `json:"closed"`
	CallbackLater  string `json:"callback_later"`
}
