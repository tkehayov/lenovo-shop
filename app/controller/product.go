package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Price        float32 `json:"price"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Series       string  `json:"series"`
	ScreenSize   string  `json:"screenSize"`
	ImagePreview string  `json:"imagePreview"`
}

func AddProduct(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var product Product
	err := decoder.Decode(&product)

	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
		return
	}
	pr := persistence.Product{Price: product.Price, Name: product.Name, ScreenSize: product.ScreenSize, ImagePreview: product.ImagePreview, Category: product.Category, Series: product.Series}

	persistence.Persist(pr)
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		log.Print(err)
	}

	pr := persistence.Get(id)
	prDom := Product{pr.Price, pr.Name, pr.Category, pr.Series, pr.ScreenSize, pr.ImagePreview}

	b := marshal(prDom)
	w.Write(b)
}

func GetAllProduct(w http.ResponseWriter, req *http.Request) {
	pr := persistence.GetAll()

	b := marshal(pr)
	w.Write(b)
}
