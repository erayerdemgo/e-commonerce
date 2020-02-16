package db

import (
	"database/sql"
	"e-commenerce/configuration"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var Dbconnection *sql.DB

func Initializedatabase() error {
	conf := configuration.Configurationins.Database
	fmt.Println("db infos", conf)
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		conf.Username, conf.Password, conf.Dbname, conf.Sslmode, conf.Host, conf.Port)
	dbconf, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return errors.New("connection problem")
	}
	Dbconnection = dbconf
	err = Dbconnection.Ping()
	if err == nil {
		fmt.Println("başarılı connection")
	}

	return nil
}
