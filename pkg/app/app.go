package app

import (
	"AvitoTst/pkg/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/deposit", handler.HomeHandler)
	r.HandleFunc("/writeOff", handler.HomeHandler)
	r.HandleFunc("/transfer", handler.HomeHandler)
	r.HandleFunc("/getBalance", handler.HomeHandler)
	//http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
