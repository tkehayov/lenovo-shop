package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lenovo-shop/app/config"
	"github.com/lenovo-shop/app/model/cart"
	"log"
)

type Order struct {
	ID        int
	Firstname string
	Lastname  string
	Address   string
	Location  string
	Email     string
	Cart      []cart.CartCookie
}

func MakeDelivery(order Order) {
	db, err := sql.Open("mysql", config.DbUri)
	tx, errT := db.Begin()
	if errT != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Println(err.Error()) // proper error handling instead of panic in your app
	}

	stmtOrderProduct, err := tx.Prepare("INSERT INTO order_product(order_id,product_id,quantity) VALUES( ?,?,? )")
	stmtOrder, err := tx.Prepare("INSERT INTO orders(firstname,lastname,address,location,email) VALUES( ?,?,?,?,? )")
	if err != nil {
		log.Print(err)
	}
	result, bulkErr := stmtOrder.Exec(order.Firstname, order.Lastname, order.Address, order.Location, order.Email)
	if bulkErr != nil {
		log.Println(bulkErr.Error())
	}

	lastId, errLastId := result.LastInsertId()

	if errLastId != nil {
		log.Println(errLastId.Error())
	}

	stmtOrderProduct.Exec(lastId, order.Cart[0].Id, order.Cart[0].Quantity)

	tx.Commit()
	fmt.Println("success")
}
