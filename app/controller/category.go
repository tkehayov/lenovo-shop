package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetCategoryProducts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]

	log.Print(category)
}
