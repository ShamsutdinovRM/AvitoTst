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

	rowUser, err := db.DB.Query("SELECT balance FROM users WHERE id=$1", body.Id)
	if err != nil {
		return model.User{}, fmt.Errorf("error check new balance: %s", err)
	}

	var newBalance model.User
	for rowUser.Next() {
		rowUser.Scan(&newBalance.Balance)
	}
	newBalance.Id = body.Id
	return newBalance, nil
}

func (db *DBModel) WriteOffMoney(body model.User) (model.User, error) {

	if body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("value less than zero: %s", body.Balance)
	}

	return model.User{}, nil
}

func (db *DBModel) TransferMoney(body model.Transfer) (model.Users, error) {
	return model.Users{}, nil
}
func (db *DBModel) GetBalanceById(body model.User) (model.User, error) {
	return model.User{}, nil
}
