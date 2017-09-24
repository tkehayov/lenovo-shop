package model

import (
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/persistence"
)

type Order struct {
	FirstName string
	LastName  string
	Address   string
	Location  string
	Email     string
	Cart      []cart.CartCookie
}

func Checkout(order Order) {
	o := persistence.Order{
		Firstname: order.FirstName,
		Lastname:  order.LastName,
		Address:   order.Address,
		Location:  order.Location,
		Email:     order.Email,
		Cart:      order.Cart,
	}

	persistence.MakeDelivery(o)

}
