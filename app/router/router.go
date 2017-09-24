package router

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/controller"
	"github.com/lenovo-shop/app/shared"
	"net/http"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	// CART
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", controller.DeleteCart).Methods("DELETE")

	// CHECKOUT
	r.HandleFunc("/checkout", controller.Checkout).Methods("POST")

	// PRODUCTS
	r.HandleFunc("/product", controller.AddProduct).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(mode.StaticPath())))
	return r
}
