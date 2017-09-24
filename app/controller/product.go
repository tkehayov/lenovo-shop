package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
)

type Product struct {
	Price float32
	Name  string
}

func AddProduct(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var product Product
	err := decoder.Decode(&product)

	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
		return
	}

	pr := persistence.ProductDell{Price: product.Price, Name: product.Name}

	persistence.Persist(pr)
}