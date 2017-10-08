package shared

import (
	"log"
	"os/user"
)

type Mode interface {
	StaticPath() string
}
type DevMode struct {
}
type ProdMode struct {
}

func (prodMode ProdMode) StaticPath() string {
	return "./static"
}

func (dev DevMode) StaticPath() string {
	return frontEndFolder()
}

func frontEndFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/projects/lenovo-shop/front-end"
}
