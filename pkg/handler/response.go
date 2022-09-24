package handler

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	s := "asdasdasdasdasd"
	sl, err := json.Marshal(s)
	if err != nil {
		return
	}

	w.Write(sl)
}

func Deposit(userDep string, b decimal.Decimal) {
}

func WriteOff(userWO string, b decimal.Decimal) {

}

func Transfer(userWO, userDep string, money decimal.Decimal) {
}

func GetBalance(user string) {

}
