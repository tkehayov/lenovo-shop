package persistence

import (
	"database/sql"
	"fmt"
	"github.com/lenovo-shop/app/model"
	"log"
	"github.com/lenovo-shop/app/config"
)

type Product struct {
}

func (p Product) Write(pr []model.Product) {
	log.Print(pr)

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

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO products VALUES( ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	stmtIns.Exec(pr[0].Name, 2.31)
	stmtIns.Exec(pr[1].Name, 2.31)
	//
	fmt.Println("success")

}
