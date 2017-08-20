package main

import (
	"flag"
	"fmt"
	"github.com/lenovo-shop/app/config"
	"github.com/lenovo-shop/app/persistence"
	"github.com/lenovo-shop/app/shared"
)

func main() {
	//pr := Exporter{}
	//pa := exporter.ReadProducts()
	ku := persistence.Product{Name: "da"}
	fmt.Print(ku)
	pr := []shared.Productor{ku}
	pr[0].Persist()

	//shared.Productor().Persist()
	//p := persistence.Product{Name: pa[0].Name}
	//p.Persist()
}
func init() {
	flag.StringVar(&config.DbUri, "dbUri", "root:titi89@/lenovo-shop", "Set to true to run the app in Dev mode. In Dev, it may panic on some errors.")

}
