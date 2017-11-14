package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/lenovo-shop/app/persistence"
	importer "github.com/lenovo-shop/importer/app/persistence"
	"io/ioutil"
	"log"
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

	//add IDs into Provider group
	var gr []importer.Groups
	for _, group := range groups.Group {
		g := importer.Groups{group.Id, group.Name}
		gr = append(gr, g)
	}

	importer.AddGroup(gr)
}
func GetAllGroups(writer http.ResponseWriter, request *http.Request) {
	grs := importer.GetAllGroups()

	b := marshal(grs)
	writer.Write(b)
}

func marshal(interf interface{}) []byte {
	b, err := json.Marshal(interf)
	if err != nil {
		log.Print(err)
	}
	return b
}
