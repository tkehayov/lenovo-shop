package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
)

type Product struct {
	Price      float32
	Name       string
	ScreenSize string
}

func AddProduct(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var product Product
	err := decoder.Decode(&product)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(400)
		return
	}

	pr := persistence.Product{Price: product.Price, Name: product.Name, ScreenSize: product.ScreenSize}

	persistence.Persist(pr)
}

func GetProduct(w http.ResponseWriter, req *http.Request) {

	persistence.Get(1)
}
func GetAllProduct(w http.ResponseWriter, req *http.Request) {
	persistence.GetAll()

}
