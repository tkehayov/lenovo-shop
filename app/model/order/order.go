package order

import (
	"github.com/lenovo-shop/app/persistence"
)

type Order struct {
	FirstName string
	LastName  string
	Address   string
	Location  string
	Email     string
}

func Checkout(order Order) {
	//firstName
	//lastName
	//street¡
	o := persistence.Order{
		Firstname: order.FirstName,
		Lastname:  order.LastName,
		Address:   order.Address,
		Location:  order.Location,
		Email:     order.Email,
	}

	persistence.MakeDelivery(o)

}
