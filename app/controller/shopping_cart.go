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
	Id       int64 `json:"id"`
	Quantity int   `json:"quantity"`
}

type ShoppingCart struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ShoppingCartCookie struct {
	ShoppingCart []ShoppingCart `json:"shoppingCarts"`
	OverallPrice float64        `json:"overallPrice"`
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
	//TODO getting ids from request
	var sc []ShoppingCart
	var ids []int64
	var overAllPrice float64

	cart, err := cart.Get(req)
	if err != nil {
		w.Write([]byte{})
		return
	}
	for _, value := range cart {
		ids = append(ids, value.Id)
	}

	pr := persistence.GetMulti(ids...)

	for index, value := range cart {
		scart := ShoppingCart{pr[index].Id, pr[index].Name, pr[index].Price, value.Quantity}
		overAllPrice = overAllPrice + (pr[index].Price * float64(value.Quantity))
		sc = append(sc, scart)
	}

	c := ShoppingCartCookie{sc, overAllPrice}

	b := marshal(c)
	w.Write(b)
}

func DeleteCart(w http.ResponseWriter, req *http.Request) {
	cart.Delete(w, req)
}

func marshal(cookie interface{}) []byte {
	b, err := json.Marshal(cookie)
	if err != nil {
		log.Print(err)
	}
	return b
}
