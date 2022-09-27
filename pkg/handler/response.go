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

func (b *Repos) Deposit(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return
	}
	defer r.Body.Close()

	var dep model.User
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		SendErr(w, http.StatusBadRequest, "Invalid field")
		return
	}

	balance, err := b.Repository.DepositMoney(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	} else {
		SendOK(w, http.StatusOK, balance)
	}

}

func (b *Repos) WriteOff(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return
	}
	defer r.Body.Close()

	var dep model.User
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		SendErr(w, http.StatusBadRequest, "Invalid field")
		return
	}

	balance, err := b.Repository.WriteOffMoney(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	} else {
		SendOK(w, http.StatusOK, balance)
	}

}

func (b *Repos) Transfer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return
	}
	defer r.Body.Close()

	var dep model.Transfer
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		SendErr(w, http.StatusBadRequest, "Invalid field")
		return
	}

	balance, err := b.Repository.TransferMoney(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	} else {
		SendOK(w, http.StatusOK, balance)
	}
}

func (b *Repos) GetBalance(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return
	}
	defer r.Body.Close()

	var dep model.Id
	if err = json.Unmarshal(body, &dep); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		SendErr(w, http.StatusBadRequest, "Invalid field")
		return
	}

	balance, err := b.Repository.GetBalanceById(dep)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
	} else {
		SendOK(w, http.StatusOK, balance)
	}
}

// api key s9kgW7F9oXVJN9lck8PoKzs4rkhb1Nf2
func (b *Repos) GetBalanceWithCurrency(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return
	}
	defer r.Body.Close()

	var req model.BalanceReq
	if err = json.Unmarshal(body, &req); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		SendErr(w, http.StatusBadRequest, "Invalid field")
		return
	}

	balance, err := b.Repository.GetBalanceById(model.Id{Id: req.Id})
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	curReturn, err := ChangeCurrency(model.BalanceCur{
		Id:       req.Id,
		Currency: req.Currency,
		Balance:  balance.Balance,
	})
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
		return
	} else {
		SendOK(w, http.StatusOK, curReturn)
	}
}

func ChangeCurrency(dep model.BalanceCur) (model.BalanceCur, error) {
	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%s&from=RUB&amount=%f", dep.Currency, dep.Balance)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		return model.BalanceCur{}, err
	}

	req.Header.Set("apikey", "s9kgW7F9oXVJN9lck8PoKzs4rkhb1Nf2")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error do request: %s\n", err)
		return model.BalanceCur{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error read body: %s\n", err)
		return model.BalanceCur{}, err
	}

	var cur model.CurResponse
	if err = json.Unmarshal(body, &cur); err != nil {
		fmt.Printf("Error unmarshal body: %s\n", err)
		return model.BalanceCur{}, err
	}

	repCur := model.BalanceCur{
		Id:       dep.Id,
		Currency: dep.Currency,
		Balance:  cur.Result,
	}
	return repCur, nil
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
