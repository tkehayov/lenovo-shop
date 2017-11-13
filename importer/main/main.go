package main

import (
	"github.com/lenovo-shop/app/shared"
	"github.com/lenovo-shop/importer/router"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	mode := shared.DevMode{}
	//production Mode
	//mode := shared.ProdMode{}

	http.Handle("/", router.GetRouter(mode))
	appengine.Main()
}
