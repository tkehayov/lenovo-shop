package controller

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
	"sort"
	"strconv"
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

	screenSizes := getMultiParam(req, "screenSizes", ",")
	priceRange := getMultiParam(req, "priceRange", ",")

	//Price From
	prices := []string{}
	for _, pricesFromTo := range priceRange { //[200-400,0-200]
		prices = append(prices, strings.Split(pricesFromTo, "-")...) // 200,400,0,200
	}

	min, max := normalizePriceRange(prices...)
	filter := persistence.Filter{ScreenSizes: screenSizes, Category: category, PriceRangeFrom: float32(min), PriceRangeTo: float32(max)}

	products := persistence.FilterProducts(filter)

	prods := []Product{}
	for _, pr := range products {
		prods = append(prods, Product{pr.Price, pr.Name, pr.ScreenSize})
	}

	b := marshal(prods)
	w.Write(b)
}

func normalizePriceRange(params ...string) (min float64, max float64) {
	data := []float64{}
	for _, param := range params {
		fl, err := strconv.ParseFloat(param, 64)
		if err != nil {
			log.Print(err)
		}
		data = append(data, fl)
	}
	sort.Float64s(data)

	return data[0], data[len(data)-1]
}

func getMultiParam(req *http.Request, key string, separator string) []string {
	params := req.URL.Query()
	screenSizesParam := params.Get(key)
	screenSizes := strings.Split(screenSizesParam, separator)
	return screenSizes
}
