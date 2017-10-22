package controller

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/persistence"
	"net/http"
	"strings"
)

type Filter struct {
	ID         int64    `json:"id"`
	PriceRange []string `json:"price"`
	Name       string   `json:"name"`
	ScreenSize string   `json:"screenSize"`
}

func FilterProducts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]

	params := req.URL.Query()
	screenSizesParam := params.Get("screenSizes")
	screenSizes := strings.Split(screenSizesParam, ",")

	filter := persistence.Filter{ScreenSizes: screenSizes, Category: category}

	products := persistence.FilterProducts(filter)

	prods := []Product{}
	for _, pr := range products {
		prods = append(prods, Product{pr.Price, pr.Name, pr.ScreenSize})
	}

	b := marshal(prods)
	w.Write(b)
}
