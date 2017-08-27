package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/persistence"
	"io/ioutil"
	"log"
	"net/http"
)

type ShoppingCart struct {
	ID       string
	Quantity int
}

func AddCart(w http.ResponseWriter, req *http.Request) {
	var sc ShoppingCart

	b, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(b, &sc)

	cs := cart.CartCookie{sc.ID, sc.Quantity}

	cc := make([]cart.CartCookie, 0)
	cart.Add(w, req, cs, &cc)

	b = marshal(cc)

	w.Write(b)
}

func GetCart(w http.ResponseWriter, req *http.Request) {
	cart, err := cart.Get(req)

	if err != nil {
		log.Fatal(err)
	}

	persistence.Get(1, 2, 3)
	b := marshal(cart)

	w.Write(b)
}

func marshal(cookie []cart.CartCookie) []byte {
	b, err := json.Marshal(cookie)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
