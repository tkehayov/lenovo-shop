package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/persistence"
	"io/ioutil"
	"log"
	"net/http"
)

type Cart struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type ShoppingCart struct {
	Name     string
	Price    float32
	Quantity int
}

func AddCart(w http.ResponseWriter, req *http.Request) {
	var sc Cart

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}
	json.Unmarshal(b, &sc)
	cs := cart.CartCookie{sc.Id, sc.Quantity}

	cc := []cart.CartCookie{}
	cart.Add(w, req, cs, &cc)

	b = marshal(cc)

	w.Write(b)
}

func GetCart(w http.ResponseWriter, req *http.Request) {
	var sc []ShoppingCart
	var ids []int

	cart, err := cart.Get(req)

	if err != nil {
		w.Write([]byte{})
		return
	}

	for _, value := range cart {
		ids = append(ids, value.ID)
	}

	pr := persistence.Get(ids...)

	for index, value := range cart {
		scart := ShoppingCart{pr[index].Name, pr[index].Price, value.Quantity}
		sc = append(sc, scart)
	}

	b := marshal(sc)

	w.Write(b)
}

func marshal(cookie interface{}) []byte {
	b, err := json.Marshal(cookie)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
