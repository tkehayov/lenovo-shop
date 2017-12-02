package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	appPers "github.com/lenovo-shop/app/persistence"
	"github.com/lenovo-shop/app/shared"
	importer "github.com/lenovo-shop/importer/app/persistence"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Group struct {
	Group []Groups `xml:"group"`
}

type Groups struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

type SubGroup struct {
	//GroupId  string      `xml:"group_id,attr"`
	Subgroup []SubGroups `xml:"subgroup"`
}

type SubGroups struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

func GetGroups(mode shared.Mode) {
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
		cat := appPers.Category{group.Name}
		appPers.AddCategory(cat)
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

func GetSubGroups(mode shared.Mode) ([]importer.SubGroups, []importer.Groups) {
	groups := importer.GetAllGroups()
	var unmSubGroups SubGroup

	for _, gr := range groups {
		resp, err := http.Get(mode.VendorUrls()["subgroups"] + gr.Id)
		log.Print("grass ", mode.VendorUrls()["subgroups"]+gr.Id)
		if err != nil {
			log.Print(err)
		}
		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			log.Print(errRead)
		}
		fmt.Println("response Body subgroups: ", string(body))
		xml.Unmarshal(body, &unmSubGroups)
	}

	var subGroupsImporter []importer.SubGroups
	for _, subGro := range unmSubGroups.Subgroup {
		subGrImp := importer.SubGroups{
			subGro.Id,
			subGro.Name,
		}
		subGroupsImporter = append(subGroupsImporter, subGrImp)
	}
	log.Print("subGroupsImporter", subGroupsImporter)
	importer.AddSubGroups(subGroupsImporter)

	return subGroupsImporter, groups
}

func generateSlug(name string) string {
	r := strings.NewReplacer(" ", "-")

	return r.Replace(name)
}

func GetAllSubGroups(groups []importer.Groups) []SubGroups {
	log.Print(groups)

	var subGroups SubGroup
	//xml.Unmarshal(b, &subGroups)

	return subGroups.Subgroup
}

func marshal(interf interface{}) []byte {
	b, err := json.Marshal(interf)
	if err != nil {
		log.Print(err)
	}
	return b
}
