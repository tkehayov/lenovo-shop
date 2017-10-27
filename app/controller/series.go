package controller

import (
	"encoding/json"
	"github.com/lenovo-shop/app/persistence"
	"log"
	"net/http"
)

type Series struct {
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func AddSeries(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var series Series
	err := decoder.Decode(&series)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(400)
		return
	}
	ser := persistence.Series{series.Name, persistence.Category{Name: series.Category.Name}}

	persistence.AddSeries(ser)
}
