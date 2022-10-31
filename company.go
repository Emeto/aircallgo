package aircallgo

type Company struct {
	Name         string `json:"name"`
	UsersCount   int64  `json:"users_count"`
	NumbersCount int64  `json:"numbers_count"`
}
