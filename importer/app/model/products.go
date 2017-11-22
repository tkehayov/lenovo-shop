package model

import (
	"encoding/xml"
	"fmt"
	app "github.com/lenovo-shop/app/persistence"
	"github.com/lenovo-shop/app/shared"
	importer "github.com/lenovo-shop/importer/app/persistence"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Products []Products `xml:"product"`
}

type Products struct {
	Description  string  `xml:"description"`
	Model        string  `xml:"model"`
	Name         string  `xml:"name"`
	Kkprice      float64 `xml:"kkprice"`
	Price        float64 `xml:"price"`
	PriceDds     float64 `xml:"price_dds"`
	SmallPicture string  `xml:"small_picture"`
}

func GetProducts(mode shared.Mode, subgr []importer.SubGroups, group []importer.Groups) {
	var pr Product
	for _, gr := range group {
		for _, s := range subgr {
			resp, err := http.Get(mode.VendorUrls()["base"] + "/products/1/" + gr.Id + "/" + s.Id)
			if err != nil {
				log.Print(err)
			}
			body, errRead := ioutil.ReadAll(resp.Body)
			if errRead != nil {
				log.Print(errRead)
			}
			fmt.Println("response Body:", string(body))
			xml.Unmarshal(body, &pr)
		}
	}

	var persProd []importer.Product
	var appProds []app.Product
	for _, prods := range pr.Products {

		products := importer.Product{
			prods.Description,
			prods.Model,
			prods.Name,
			prods.Kkprice,
			prods.Price,
			prods.PriceDds,
			prods.SmallPicture,
		}

		// TODO Baby
		appProd := app.Product{}

		persProd = append(persProd, products)
		appProds = append(appProds, appProd)

		log.Print("persProd", persProd)
	}

	log.Print("appProds", appProds)

	app.PersistMulti(appProds)
	importer.AddMultiProduct()
	log.Print(pr.Products)
}
