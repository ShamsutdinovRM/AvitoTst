package handler

import (
	"AvitoTst/model"
	"AvitoTst/pkg/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repos struct {
	Repository repository.Operations
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	s := "asdasdasdasdasd"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

func (b *Repos) Deposit(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s", err)
	}
	defer r.Body.Close()

	var dep model.User
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s", err)
	}

	balance, err := b.Repository.DepositMoney(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	}

	SendOK(w, http.StatusOK, balance)
}

func (b *Repos) WriteOff(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s", err)
	}
	defer r.Body.Close()

	var dep model.User
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s", err)
	}

	balance, err := b.Repository.WriteOffMoney(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	}

	SendOK(w, http.StatusOK, balance)
}

func (b *Repos) Transfer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s", err)
	}
	defer r.Body.Close()

	var dep model.Transfer
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s", err)
	}
}

func (b *Repos) GetBalance(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s", err)
	}
	defer r.Body.Close()

	var dep model.User
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s", err)
	}
}

func SendErr(w http.ResponseWriter, code int, text string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(
		model.DefaultError{Text: text},
	)
}

func SendOK(w http.ResponseWriter, code int, p interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		p,
	)
}
