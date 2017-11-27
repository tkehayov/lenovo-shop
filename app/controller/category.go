package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
)

type Category struct {
	Name string `json:"name"`
}

func AddCategory(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var category Category
	err := decoder.Decode(&category)

	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
		return
	}

	persistence.AddCategory(persistence.Category{Name: category.Name})
}

func GetAllCategories(w http.ResponseWriter, req *http.Request) {
	cats := persistence.GetAllCategories()

	b := marshal(cats)
	w.Write(b)
}
