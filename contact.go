package aircallgo

type Contact struct {
	ID           int64                `json:"id"`
	DirectLink   string               `json:"direct_link"`
	FirstName    string               `json:"first_name"`
	LastName     string               `json:"last_name"`
	CompanyName  string               `json:"company_name"`
	Description  string               `json:"description"`
	Information  string               `json:"information"`
	IsShared     bool                 `json:"is_shared"`
	PhoneNumbers []ContactPhoneNumber `json:"phone_numbers"`
	Emails       []ContactEmail       `json:"emails"`
}

type ContactPhoneNumber struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type ContactEmail struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}
