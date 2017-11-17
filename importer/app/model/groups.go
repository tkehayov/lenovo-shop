package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/lenovo-shop/app/persistence"
	"github.com/lenovo-shop/app/shared"
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

type SubGroup struct {
	GroupId  string      `xml:"group_id,attr"`
	Subgroup []SubGroups `xml:"subgroup"`
}

type SubGroups struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

func GetGroups(mode shared.Mode) {
	//vars := mux.Vars(request)
	//id := vars["id"]
	resp, err := http.Get(mode.VendorUrls()["groups"])
	if err != nil {
		log.Print(err)
	}

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Print(errRead)
	}
	fmt.Println("response Body:", string(body))

	var groups Group
	xml.Unmarshal(body, &groups)

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

func GetSubGroups(writer http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//vendorId := vars["vendorId"]

	groups := importer.GetAllGroups()
	subgr := GetAllSubGroups(groups)

	log.Print(subgr)
}

func GetAllSubGroups(groups []importer.Groups) []SubGroups {

	xmlFile, err := os.Open("subgroups.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var subGroups SubGroup
	xml.Unmarshal(b, &subGroups)

	return subGroups.Subgroup
}

func marshal(interf interface{}) []byte {
	b, err := json.Marshal(interf)
	if err != nil {
		log.Print(err)
	}
	return b
}
