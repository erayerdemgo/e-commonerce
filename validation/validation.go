package validation

import (
	"e-commenerce/db"
	"fmt"
	"log"

	"github.com/go-playground/validator"
)

var Validate = validator.New()

const (
	existmail = "SELECT  (1=1) from customer where email=$1"
)

func Existmail(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	prepare, err := db.Dbconnection.Prepare(existmail)
	defer prepare.Close()
	if err != nil {
		log.Fatalf("problem %v", err)
	}
	exist := false
	rows, err := prepare.Query(s)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&exist); err != nil {
			fmt.Println("rows", err)
		}
	}
	fmt.Println(exist, "value")
	return !exist
}

func Createcustomvalidations() error {

	if err := Validate.RegisterValidation("checkdb", Existmail); err != nil {
		return err

	}
	return nil
}
