package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type Filter struct {
	ScreenSizes []string
	Category    string
}

func FilterProducts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]

	params := req.URL.Query()
	screenSizesParam := params.Get("screenSizes")
	screenSizes := strings.Split(screenSizesParam, ",")

	cat := Filter{ScreenSizes: screenSizes, Category: category}
	log.Print("cat", cat)
	log.Print("category", category)
	log.Print("params", screenSizes)
}
