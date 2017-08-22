package controller

import (
	"encoding/json"
	"fmt"
	"github.com/lenovo-shop/app/model/cart"
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

	b, err := json.Marshal(cc)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)
}

func GetCart(w http.ResponseWriter, req *http.Request) {
	cookie, err := cart.Get(req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(cookie)
}
