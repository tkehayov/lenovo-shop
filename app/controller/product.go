package controller

import (
	"github.com/lenovo-shop/app/persistence"
	"net/http"
)

func AddProduct(w http.ResponseWriter, req *http.Request) {

	//Get Params from request
	pr := persistence.ProductDell{Price: 23.2, Name: "Chushki"}

	persistence.Persist(pr)
}
