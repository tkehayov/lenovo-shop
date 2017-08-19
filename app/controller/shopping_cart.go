package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/model/cart"
	"io/ioutil"
	"log"
	"net/http"
)

type ShoppingCart struct {
	ID          string
	ProductName string
	Quantity    int
}

func AddCart(w http.ResponseWriter, req *http.Request) {
	var sc ShoppingCart

	b, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(b, &sc)

	cs := cart.Cookie{sc.ID, sc.ProductName, sc.Quantity}
	cart.Add(w, cs)

	w.WriteHeader(201)
}

func GetCart(w http.ResponseWriter, req *http.Request) {
	cookie, err := cart.Get(req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(cookie.ProductName)
}
