package main

import (
	"github.com/lenovo-shop/app/router"
	"github.com/lenovo-shop/app/shared"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	mode := shared.DevMode{}
	//Uncomment for production
	//mode:=shared.ProdMode{}

	http.Handle("/", router.GetRouter(mode))
	appengine.Main()
}