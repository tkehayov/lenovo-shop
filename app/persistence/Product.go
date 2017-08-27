package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lenovo-shop/app/config"
	"log"
	"strings"
)

type Product struct {
	Name string
}

func Persist(pr Product) {
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

	stmtIns, err := db.Prepare("INSERT INTO products VALUES( ?, ? )")
	if err != nil {
		log.Print(err)
	}
	defer stmtIns.Close()

	stmtIns.Exec(pr.Name, 2.31)

	fmt.Println("success")
}

func Get(ids ...int) {
	db, err := sql.Open("mysql", config.DbUri)
	if err != nil {

		log.Fatal(err.Error())
	}

	args := []interface{}{}
	for _, value := range ids {
		args = append(args, value)
	}

	selDB, err := db.Query("SELECT id FROM products WHERE id in (?"+strings.Repeat(",?", len(ids)-1)+")", args...)

	if err != nil {
		log.Fatal(err.Error())
	}

	for selDB.Next() {
		var id string

		err = selDB.Scan(&id)

		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
}
