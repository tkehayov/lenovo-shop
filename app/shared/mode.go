package shared

import (
	"log"
	"os/user"
)

type Mode interface {
	StaticPath() string
	ImagePath() string
	VendorUrls() map[string]string
}

type DevMode struct {
}

type ProdMode struct {
}

func (prodMode ProdMode) StaticPath() string {
	return "./static"
}

//TODO change prodmode path
func (dev ProdMode) ImagePath() string {
	return frontEndFolder() + "/product-images"
}

func (dev DevMode) StaticPath() string {
	return frontEndFolder()
}

func (dev DevMode) ImagePath() string {
	return "../product-images/"
}

func (dev DevMode) VendorUrls() map[string]string {
	baseUrl := "http://localhost:9000"

	return map[string]string{
		"groups":    baseUrl + "/groups/1",
		"subgroups": baseUrl + "/subgroups/{vendorId}/{groupId}",
	}
}

func (prodMode ProdMode) VendorUrls() map[string]string {
	return map[string]string{
		"rsc": "",
	}
}

func frontEndFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Print(err)
	}

	return usr.HomeDir + "/projects/lenovo-shop/front-end"
}
