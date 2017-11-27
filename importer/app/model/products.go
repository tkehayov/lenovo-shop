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
	GroupdId string     `xml:"groupd_id,attr"`
	Products []Products `xml:"product"`
}

type Products struct {
	Id string `xml:"id"`
}

type ProductId struct {
	Id string
}

type AllProducts struct {
	//ProductInfo ProductInfo `xml:"product_info"`
	Group           string          `xml:"group"`
	Model           string          `xml:"model"`
	Characteristics Characteristics `xml:"characteristics"`
	Warranty        int16           `xml:"months"`
	Price           float64         `xml:"price"`
	PriceVat        float64         `xml:"price_dds"`
	SubGroups       string          `xml:"subgroup"`
}

type Characteristics struct {
	Characteristic []Characteristic `xml:"characteristic"`
}

type Characteristic struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

func GetProducts(mode shared.Mode, subgr []importer.SubGroups, group []importer.Groups) []ProductId {
	var pr Product
	var productIds []ProductId

	var prodsId []importer.ProductId
	for _, gr := range group {
		for _, s := range subgr {
			resp, err := http.Get(mode.VendorUrls()["products"] + gr.Id + "/" + s.Id)
			if err != nil {
				log.Print(err)
			}
			body, errRead := ioutil.ReadAll(resp.Body)
			if errRead != nil {
				log.Print(errRead)
			}

			fmt.Println("response Body products:", string(body))
			xml.Unmarshal(body, &pr)
		}
	}

	for _, prods := range pr.Products {
		prodId := importer.ProductId{
			prods.Id,
		}
		prodsId = append(prodsId, prodId)
	}

	importer.PersistMultiProducts(prodsId...)
	prods := importer.GetAllProducts()

	for _, prod := range prods {
		productId := ProductId{
			prod.Id,
		}
		productIds = append(productIds, productId)
	}

	return productIds
}

func GetAllProducts(prods []ProductId, mode shared.Mode) {
	var allProducts []AllProducts
	var products []app.Product

	for _, prod := range prods {
		resp, err := http.Get(mode.VendorUrls()["product"] + prod.Id)

		if err != nil {
			log.Print(err)
		}

		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			log.Print(errRead)
		}
		log.Print("response Body all products ", string(body))

		xml.Unmarshal(body, &allProducts)
	}

	for _, product := range allProducts {
		pr := app.Product{
			Price:       product.Price,
			Category:    product.Group,
			SubCategory: product.SubGroups,
			PriceVat:    product.PriceVat,
			Warranty:    product.Warranty,
			Model:       product.Model,
		}
		products = append(products, pr)
	}

	app.PersistMulti(products)
	log.Print("allProducts", allProducts)
}
