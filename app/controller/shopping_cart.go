package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/model/cart"
	"github.com/lenovo-shop/app/model/order"
	//"github.com/lenovo-shop/app/persistence"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Cart struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type ShoppingCart struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ShoppingCartCookie struct {
	ShoppingCart []ShoppingCart `json:"shoppingCarts"`
	OverallPrice float32        `json:"overallPrice"`
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
	log.Println("Welcome on cart")
	var sc []ShoppingCart
	var ids []int
	var overAllPrice float32

	cart, err := cart.Get(req)

	if err != nil {
		w.Write([]byte{})
		return
	}
	for _, value := range cart {
		ids = append(ids, value.Id)
	}

	//pr := persistence.Get(ids...)

	//for index, value := range cart {
	//scart := ShoppingCart{pr[index].ID, pr[index].Name, pr[index].Price, value.Quantity}
	//overAllPrice = overAllPrice + (pr[index].Price * float32(value.Quantity))
	//sc = append(sc, scart)
	//}

	c := ShoppingCartCookie{sc, overAllPrice}

	b := marshal(c)
	w.Write(b)
}

func DeleteCart(w http.ResponseWriter, req *http.Request) {
	cart.Delete(w, req)
}

func Checkout(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	firstName := req.Form["firstName"][0]
	lastName := req.Form["lastName"][0]
	address := req.Form["address"][0]
	location := req.Form["location"][0]
	email := req.Form["email"][0]

	cookies, errCookie := cart.Get(req)

	if errCookie != nil {
		//TODO redirect to homepage
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(10 * time.Second)
			timeout <- true
			http.Redirect(w, req, "http://www.google.com", 301)

			return
		}()

		return
	}

	d := order.Order{firstName, lastName, address, location, email, cookies}
	order.Checkout(d)
}

func marshal(cookie interface{}) []byte {
	b, err := json.Marshal(cookie)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
