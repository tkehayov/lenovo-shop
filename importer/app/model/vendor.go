package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Vendor struct {
	Vendor []Vendors `xml:"vendor"`
}

type Vendors struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

func GetVendor(writer http.ResponseWriter, request *http.Request) {
	xmlFile, err := os.Open("vendor.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	var vendor Vendor
	xml.Unmarshal(b, &vendor)

	log.Println(vendor.Vendor)
}
