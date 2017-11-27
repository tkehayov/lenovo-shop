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
