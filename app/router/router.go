package router

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/controller"
	"github.com/lenovo-shop/app/shared"
	"net/http"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	// Cart
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", controller.DeleteCart).Methods("DELETE")

	// Orders
	r.HandleFunc("/order", controller.Order).Methods("POST")
	r.HandleFunc("/order", controller.ListOrders).Methods("GET")

	// Products
	r.HandleFunc("/product", controller.AddProduct).Methods("POST")
	// TODO add some params in endpoint
	r.HandleFunc("/product", controller.GetProduct).Methods("GET")
	r.HandleFunc("/products/all", controller.GetAllProduct).Methods("GET")

	//Categories
	r.HandleFunc("/categories/{category}", controller.FilterProducts).Methods("GET")
	r.HandleFunc("/categories", controller.AddCategory).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(mode.StaticPath())))

	return r
}
