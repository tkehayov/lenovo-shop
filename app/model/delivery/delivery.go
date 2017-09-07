package delivery

import (
	"fmt"
)

type Delivery struct {
	FirstName string
	LastName  string
}

func Checkout(delivery Delivery) {
	//firstName
	//lastName
	//street

	fmt.Print(delivery)
}
