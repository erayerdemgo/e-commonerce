package service

import (
	"e-commenerce/db"
	"e-commenerce/model"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	insertuser = "INSERT INTO customer(uuid, name, surname, email, password)\nVALUES ($1, $2, $3, $4, $5);"
)

func Createuser(user model.UserSignup) error {
	db := db.Dbconnection
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	//defer tx.Rollback()
	stmt, err := tx.Prepare(insertuser)
	if err != nil {
		log.Fatal(err)
	}
	// danger!
	password := []byte(user.Password)
	hashedpassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hash problem")
	}

	_, err = stmt.Exec(uuid.New(), user.Name, user.Surname, user.Email, hashedpassword)
	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err, "commit problem")
	}
	return nil
}

func Login(user model.UserSignin) bool {
	fmt.Println("senin user",user)

	prepare, err := db.Dbconnection.Prepare("select password from customer where email=$1")
	if err != nil {
		fmt.Println(err)
	}
	query, err := prepare.Query(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	hashedpass := ""
	for query.Next() {
		err := query.Scan(&hashedpass)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer query.Close()
	prepare.Close()
	fmt.Println(hashedpass)
	err = bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte (user.Password))
	if err == nil {

		fmt.Println("hak ettiniz tebrikler")
		return true
	}
	return false
}
