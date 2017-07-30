package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Exporter struct {
	Category []Category `json:"categories"`
}
type Category struct {
	Name string `json:"name"`
}

func ReadExport(filename string) {
	res := Exporter{}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	enmErr := json.Unmarshal(content, &res)
	if enmErr != nil {
		log.Fatal(enmErr)
	}

	for cats := range res.Category {
		fmt.Println(cats)

	}
}
