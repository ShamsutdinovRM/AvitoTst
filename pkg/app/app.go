package app

import (
	"AvitoTst/model"
	"AvitoTst/pkg/handler"
	"AvitoTst/pkg/repository"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Run(path string) {
	if err := initConfig(path); err != nil {
		fmt.Printf("error initializaing configs: %s", err)
		return
	}

	DBSchema := model.DB{
		Username: viper.GetString("db.username"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Schema:   viper.GetString("db.schema"),
		Password: viper.GetString("db.password"),
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DBSchema.Username, DBSchema.Password, DBSchema.DBName, DBSchema.SSLMode)

	db, err := repository.New(connStr)
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

	port := viper.GetString("port")
	fmt.Printf("server started")
	log.Fatal(http.ListenAndServe(port, r))
}

func initConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
