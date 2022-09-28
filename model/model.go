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

type BalanceCur struct {
	Id       string  `json:"id"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type BalanceReq struct {
	Id       string `json:"id"`
	Currency string `json:"currency"`
}

type CurResponse struct {
	Success bool    `json:"success"`
	Query   Query   `json:"query"`
	Info    Info    `json:"info"`
	Date    string  `json:"date"`
	Result  float64 `json:"result"`
}

type Query struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type Info struct {
	Timestamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
}

type Id struct {
	Id string `json:"id"`
}

type DB struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	DBSchema string `json:"dbschema"`
	Password string `json:"password"`
}

type DefaultError struct {
	Text string `json:"text"`
}
