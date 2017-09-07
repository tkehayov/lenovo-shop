package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/config"
	"github.com/lenovo-shop/app/controller"
	"net/http"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", controller.DeleteCart).Methods("DELETE")

	r.HandleFunc("/checkout", controller.Checkout).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticFolder)))
	return r
}
