package persistence

import (
	"fmt"
	"github.com/lenovo-shop/app/model/cart"
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

	fmt.Println("success")
}
