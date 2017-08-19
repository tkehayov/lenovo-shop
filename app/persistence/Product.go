package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lenovo-shop/app/config"
	"log"
)

type Product struct {
	Name string
}

func (pr Product) Persist() {
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

	stmtIns.Exec(pr.Name, 2.31)

	fmt.Println("success")
}
