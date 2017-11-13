package router

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/shared"
	"github.com/lenovo-shop/importer/model"
	"net/http"
	"reflect"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	//Vendors
	t := reflect.TypeOf(mode).String()
	if t == "shared.DevMode" {
		r.HandleFunc("/test", model.GetVendor).Methods("GET")

		return r
	}

	r.HandleFunc("/test", model.GetVendor).Methods("POST")

	return r
}
