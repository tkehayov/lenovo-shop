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

func (dev ProdMode) ImagePath() string {
	return "./static/product-images"
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
		"base":      baseUrl,
		"groups":    baseUrl + "/groups/0006301183441",
		"subgroups": baseUrl + "/subgroups/0006301183441/",
		"products":  baseUrl + "/products/0006301183441/",
		"product":   baseUrl + "/product/",
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
