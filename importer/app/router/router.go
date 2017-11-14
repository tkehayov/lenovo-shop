package router

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/shared"
	"github.com/lenovo-shop/importer/app/model"
	"net/http"
	"reflect"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	//Vendors
	t := reflect.TypeOf(mode).String()
	if t == "shared.DevMode" {
		r.HandleFunc("/test", model.GetVendor).Methods("GET")
		r.HandleFunc("/groups/{id}", model.GetGroups).Methods("GET")

		r.HandleFunc("/groupss/all", model.GetAllGroups).Methods("GET")

		return r
	}

	r.HandleFunc("/test", model.GetVendor).Methods("POST")

	return r
}
