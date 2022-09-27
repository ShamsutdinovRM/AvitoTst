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
	GetBalanceById(body model.Id) (model.User, error)
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
		return model.User{}, fmt.Errorf("value less than zero: %f", body.Balance)
	}

	actual, err := db.GetBalanceById(model.Id{Id: body.Id})
	if err != nil {
		return model.User{}, err
	}

	_, err = db.DB.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", body.Balance, body.Id)
	if err != nil {
		return model.User{}, fmt.Errorf("error deposit money: %s", err)
	}

	actual, err = db.GetBalanceById(model.Id{Id: body.Id})
	if err != nil {
		return model.User{}, err
	}

	return actual, nil
}

func (db *DBModel) WriteOffMoney(body model.User) (model.User, error) {

	if body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("value less than zero: %f", body.Balance)
	}

	actual, err := db.GetBalanceById(model.Id{Id: body.Id})
	if err != nil {
		return model.User{}, err
	}

	if actual.Balance-body.Balance < 0.0 {
		return model.User{}, fmt.Errorf("insufficient funds in the account, current balance: %f", actual.Balance)
	}

	_, err = db.DB.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", body.Balance, body.Id)
	if err != nil {
		return model.User{}, fmt.Errorf("error write off money: %s", err)
	}

	actual, err = db.GetBalanceById(model.Id{Id: body.Id})
	if err != nil {
		return model.User{}, err
	}

	return actual, nil
}

func (db *DBModel) TransferMoney(body model.Transfer) (model.Users, error) {

	var actualUserWO model.User
	actualUserWO, err := db.GetBalanceById(model.Id{Id: body.UserWO})
	if err != nil {
		return model.Users{}, err
	}

	if actualUserWO.Balance < body.Balance {
		return model.Users{}, fmt.Errorf("insufficient funds in the write off account, check your balance")
	}

	_, err = db.WriteOffMoney(model.User{Id: body.UserWO, Balance: body.Balance})
	if err != nil {
		return model.Users{}, err
	}

	_, err = db.DepositMoney(model.User{Id: body.UserDep, Balance: body.Balance})
	if err != nil {
		return model.Users{}, err
	}

	return model.Users{UserWO: body.UserWO, UserDep: body.UserDep, Status: "Ok"}, nil
}

func (db *DBModel) GetBalanceById(body model.Id) (model.User, error) {

	rowUser := db.DB.QueryRow("SELECT balance FROM users WHERE id=$1 AND balance IS NOT NULL", body.Id)

	var currentBalance model.User
	err := rowUser.Scan(&currentBalance.Balance)
	if err != nil {
		if err != sql.ErrNoRows {
			return model.User{}, fmt.Errorf("error take balance: %s", err)
		}

		_, err = db.DB.Exec("UPDATE users SET balance = 0 WHERE id = $1", body.Id)
		if err != nil {
			return model.User{}, fmt.Errorf("error create balance: %s", err)
		}
	}

	rowUser = db.DB.QueryRow("SELECT balance FROM users WHERE id=$1", body.Id)

	err = rowUser.Scan(&currentBalance.Balance)
	if err != nil {
		return model.User{}, fmt.Errorf("error take balance: %s", err)
	}
	currentBalance.Id = body.Id
	return currentBalance, nil
}
