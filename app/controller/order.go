package controller

import (
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/model/order"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
)

func Order(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	firstName := req.Form["firstName"][0]
	lastName := req.Form["lastName"][0]
	address := req.Form["address"][0]
	location := req.Form["location"][0]
	email := req.Form["email"][0]

	cookies, errCookie := cart.Get(req)

	if errCookie != nil {
		log.Fatal("error Cookie", errCookie)
		//TODO redirect to homepage
		//timeout := make(chan bool, 1)
		//go func() {
		//	time.Sleep(10 * time.Second)
		//	timeout <- true
		//	http.Redirect(w, req, "http://www.google.com", 301)
		//
		//	return
		//}()

		return
	}

	d := order.Order{firstName, lastName, address, location, email, cookies}
	order.Checkout(d)
}

func ListOrders(w http.ResponseWriter, req *http.Request) {
	persistence.ListOrders()

}
