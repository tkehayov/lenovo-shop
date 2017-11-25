package model

import (
	"encoding/xml"
	"fmt"
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
	Description  string  `xml:"description"`
	Id           string  `xml:"id"`
	Model        string  `xml:"model"`
	Name         string  `xml:"name"`
	Kkprice      float64 `xml:"kkprice"`
	Price        float64 `xml:"price"`
	PriceDds     float64 `xml:"price_dds"`
	SmallPicture string  `xml:"small_picture"`
}

type ProductId struct {
	Id string
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

func GetAllProducts(prods []ProductId) {
	log.Print("prods ", prods)

}
