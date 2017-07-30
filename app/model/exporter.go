package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Exporter struct {
	Category []Category `json:"categories"`
	Product  []Product  `json:"products"`
}

type Category struct {
	Name  string `json:"name"`
	CatId int    `json:"category_id"`
}

//Products
type Product struct {
	Manufacturer string `json:"manufacturer"`
	Warranty     string `json:"gar_srok_21"`
	CatId        string `json:"category_ids"`
	Name         string `json:"name"`
}

func ReadCategories(filename string) {
	res := Exporter{}

	content := readFile(filename)
	enmErr := json.Unmarshal(content, &res)
	if enmErr != nil {
		log.Fatal(enmErr)
	}

	var catIds int
	for i := range res.Category {
		cat := res.Category[i]
		if cat.Name == "Лаптопи" {
			catIds = cat.CatId
		}

		for p := range res.Product {
			singlep := res.Product[p]
			cats := strings.Split(singlep.CatId, ",")

			//log.Print()
			for singlecatindex := range cats {
				m, era := strconv.Atoi(cats[singlecatindex])
				if m == catIds {
					log.Println(singlep.Name)
					log.Println(singlep.Manufacturer)
					log.Println(singlep.Warranty)
					log.Println("-------")
				}

				if era != nil {
					log.Println(era)
				}
			}
		}
	}

	log.Println(res.Product[0].Manufacturer)
}

func ReadProducts() []Product {
	res := Exporter{}
	prod := []Product{}
	content := readFile("./export.json")
	enmErr := json.Unmarshal(content, &res)
	if enmErr != nil {
		log.Fatal(enmErr)
	}

	for index := range res.Product {
		p := res.Product[index]

		if p.Manufacturer == "lenovo" {
			prod = append(prod, p)
		}
	}

	return prod
}

func readFile(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	log.Println(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
