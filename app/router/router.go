package router

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/controller"
	"net/http"
	"github.com/lenovo-shop/app/shared"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", controller.DeleteCart).Methods("DELETE")

	r.HandleFunc("/checkout", controller.Checkout).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(mode.StaticPath())))
	return r
}
