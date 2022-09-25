package repository

import (
	"AvitoTst/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/shopspring/decimal"
)

type Operations interface {
	DepositMoney(body model.User) (model.User, error)
	WriteOffMoney(body model.User) (model.User, error)
	TransferMoney(body model.Transfer) (model.Users, error)
	GetBalanceById(body model.User) (model.User, error)
}

type DBModel struct {
	DB *sql.DB
}

func New() (*DBModel, error) {
	connStr := "user=dev password=dev dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error open DB: %s", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("Error Ping DB: %s", err)
		return nil, err
	}
	return &DBModel{DB: db}, nil
}

func (db *DBModel) DepositMoney(body model.User) (model.User, error) {

	if body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("value less than zero: %s", body.Balance)
	}
	_, err := db.DB.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", body.Balance, body.Id)
	if err != nil {
		return model.User{}, fmt.Errorf("error deposit money: %s", err)
	}

	actual, err := db.GetBalanceById(body)
	if err != nil {
		return model.User{}, err
	}

	return actual, nil
}

func (db *DBModel) WriteOffMoney(body model.User) (model.User, error) {

	if body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("value less than zero: %s", body.Balance)
	}

	actual, err := db.GetBalanceById(body)
	if err != nil {
		return model.User{}, err
	}

	if actual.Balance-body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("insufficient funds in the account, current balance: %s", actual.Balance)
	}

	_, err = db.DB.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", body.Balance, body.Id)
	if err != nil {
		return model.User{}, fmt.Errorf("error write off money: %s", err)
	}

	actual, err = db.GetBalanceById(body)
	if err != nil {
		return model.User{}, err
	}

	return actual, nil
}

func (db *DBModel) TransferMoney(body model.Transfer) (model.Users, error) {
	return model.Users{}, nil
}

func (db *DBModel) GetBalanceById(body model.User) (model.User, error) {

	rowUser := db.DB.QueryRow("SELECT balance FROM users WHERE id=$1", body.Id)

	var currentBalance model.User
	err := rowUser.Scan(&currentBalance.Id, &currentBalance.Balance)
	if err != nil {
		return model.User{}, fmt.Errorf("error take balance: %s", err)
	}

	return currentBalance, nil
}
