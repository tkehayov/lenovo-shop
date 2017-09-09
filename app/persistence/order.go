package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lenovo-shop/app/config"
	"log"
)

type Order struct {
	ID        int
	Firstname string
	Lastname  string
	Address   string
	Location  string
	Email     string
}

func MakeDelivery(order Order) {
	db, err := sql.Open("mysql", config.DbUri)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmtIns, err := db.Prepare("INSERT INTO orders(firstName,lastName,address,location,email) VALUES( ?,?,?,?,? )")
	if err != nil {
		log.Println("turutututu")
		log.Print(err)
	}
	defer stmtIns.Close()

	stmtIns.Exec(order.Firstname, order.Lastname, order.Address, order.Location, order.Email)

	fmt.Println("success")
}
