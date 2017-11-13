package router

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/shared"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

type Vendor struct {
	Vendor []Vendors `xml:"vendor"`
}

type Vendors struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	//Vendors
	t := reflect.TypeOf(mode).String()
	if t == "shared.DevMode" {
		r.HandleFunc("/test", GetVendor).Methods("GET")

		return r
	}

	r.HandleFunc("/test", GetVendor).Methods("POST")

	return r
}

func GetVendor(writer http.ResponseWriter, request *http.Request) {
	xmlFile, err := os.Open("vendor.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	//sss
	b, _ := ioutil.ReadAll(xmlFile)

	var vendor Vendor
	xml.Unmarshal(b, &vendor)
	//ssss
	log.Println(vendor.Vendor)
}
