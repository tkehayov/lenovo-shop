package main

import (
	"github.com/lenovo-shop/app/shared"
	"github.com/lenovo-shop/importer/app/model"
	"github.com/lenovo-shop/importer/app/router"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	mode := shared.DevMode{}
	//production Mode
	//mode := shared.ProdMode{}

	model.GetGroups(mode)
	subgr, gr := model.GetSubGroups(mode)
	prods := model.GetProducts(mode, subgr, gr)
	model.GetAllProducts(prods)

	http.Handle("/", router.GetRouter(mode))
	appengine.Main()
}
