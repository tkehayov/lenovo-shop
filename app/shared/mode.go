package shared

import (
	"log"
	"os/user"
)

type Mode interface {
	StaticPath() string
	ImagePath() string
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

func frontEndFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Print(err)
	}

	return usr.HomeDir + "/projects/lenovo-shop/front-end"
}
