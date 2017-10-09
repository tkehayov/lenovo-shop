package controller

import (
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/model/order"
	"github.com/lenovo-shop/app/persistence"
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
		http.Redirect(w, req, "/", 301)
	}

	d := order.Order{firstName, lastName, address, location, email, cookies}
	order.Checkout(d)

	http.Redirect(w, req, "/?message=Благодарим за поръчката! Ще се свържем с Вас възможно най скоро.", 301)

	return
}

func ListOrders(w http.ResponseWriter, req *http.Request) {
	persistence.ListOrders()

}
