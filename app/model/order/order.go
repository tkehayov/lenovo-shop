package order

import (
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/persistence"
)

//TODO to rename this one
type Order struct {
	FirstName string
	LastName  string
	Address   string
	Location  string
	Email     string
	Cart      []cart.CartCookie
}

//TODO move it into controller and rename to Order
func Checkout(order Order) {
	o := persistence.Order{
		Firstname: order.FirstName,
		Lastname:  order.LastName,
		Address:   order.Address,
		Location:  order.Location,
		Email:     order.Email,
	}
	persistence.MakeOrder(o, 1)

}
