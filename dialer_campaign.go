package aircallgo

import "time"

type DialerCampaign struct {
	ID        int64     `json:"id"`
	NumberID  string    `json:"number_id"`
	CreatedAt time.Time `json:"created_at"`
}

type DialerCampaignPhoneNumber struct {
	ID        int64     `json:"id"`
	Number    string    `json:"number"`
	Called    bool      `json:"called"`
	CreatedAt time.Time `json:"created_at"`
}
