package model

import (
	"encoding/xml"
	"fmt"
	"github.com/lenovo-shop/app/persistence"
	"io/ioutil"
	"net/http"
	"os"
)

type Group struct {
	Group []Groups `xml:"group"`
}
type Groups struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

func GetGroups(writer http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//id := vars["id"]

	xmlFile, err := os.Open("groups.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	var groups Group
	xml.Unmarshal(b, &groups)

	for _, group := range groups.Group {
		cat := persistence.Category{group.Name}
		persistence.AddCategory(cat)
	}

}
