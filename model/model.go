package model

type User struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Transfer struct {
	UserWO  string  `json:"userwo"`
	UserDep string  `json:"userdep"`
	Balance float64 `json:"balance"`
}

type Users struct {
	UserWO  string `json:"userwo"`
	UserDep string `json:"userdep"`
	Status  string `json:"status"`
}

type DefaultError struct {
	Text string `json:"text"`
}
