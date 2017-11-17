package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func GetProducts(writer http.ResponseWriter, request *http.Request) {
	xmlFile, err := os.Open("products.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	var product Product
	xml.Unmarshal(b, &product)

	log.Print(product)
}
