package model

type User struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Transfer struct {
	UserWO  string `json:"userwo"`
	UserDep string `json:"userdep"`
	Balance string `json:"balance"`
}

type Users struct {
	UserWO  string `json:"userwo"`
	UserDep string `json:"userdep"`
}

type DefaultError struct {
	Text string `json:"text"`
}
