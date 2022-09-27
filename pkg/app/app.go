package app

import (
	"AvitoTst/pkg/handler"
	"AvitoTst/pkg/repository"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	db, err := repository.New()
	if err != nil {
		fmt.Printf("Error create DB connection: %s", err)
		return
	}
	defer db.DB.Close()

	hand := handler.Repos{Repository: db}

	r := mux.NewRouter()
	r.HandleFunc("/deposit", hand.Deposit)
	r.HandleFunc("/writeOff", hand.WriteOff)
	r.HandleFunc("/transfer", hand.Transfer)
	r.HandleFunc("/getBalance", hand.GetBalance)
	r.HandleFunc("/getBalanceCurrency", hand.GetBalanceWithCurrency)
	log.Fatal(http.ListenAndServe(":8080", r))
}
