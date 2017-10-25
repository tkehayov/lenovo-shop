package controller

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/persistence"
	"github.com/lenovo-shop/app/shared"
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

type FilterProduct struct {
	Id           int64   `json:"id"`
	Price        float32 `json:"price"`
	Name         string  `json:"name"`
	ScreenSize   string  `json:"screenSize"`
	ImagePreview string  `json:"imagePreview"`
}

func FilterProducts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]

	mode := context.Get(req, "mode").(shared.Mode)

	screenSizes := getMultiParam(req, "screenSizes", ",")
	priceRange := getMultiParam(req, "priceRange", ",")
	log.Print(screenSizes)
	//Price From
	prices := []string{}
	for _, pricesFromTo := range priceRange { //[200-400,0-200]
		prices = append(prices, strings.Split(pricesFromTo, "-")...) // 200,400,0,200
	}

	min, max := normalizePriceRange(prices...)
	filter := persistence.Filter{ScreenSizes: screenSizes, Category: category, PriceFrom: float32(min), PriceTo: float32(max)}

	products := persistence.FilterProducts(filter)

	prods := []FilterProduct{}
	for _, pr := range products {
		log.Print("pr.Id", pr.Id)
		prods = append(prods, FilterProduct{pr.Id, pr.Price, pr.Name, pr.ScreenSize, mode.ImagePath() + pr.ImagePreview})
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
