package main

import (
	"flag"
	"github.com/lenovo-shop/app/config"
	"github.com/lenovo-shop/app/router"
	"log"
	"net/http"
	"os/user"
)

func main() {
	http.Handle("/", router.GetRouter())
	http.ListenAndServe(config.Port, nil)
}

func init() {
	log.Println("Starting...")

	flag.StringVar(&config.Port, "port", ":8080", "Specify the port to listen to e.g. :8080")
	flag.BoolVar(&config.IsDev, "isDev", true, "Set to true to run the app in Dev mode. In Dev, it may panic on some errors.")
	flag.StringVar(&config.StaticFolder, "static", frontEndFolder(), "Set to true to run the app in Dev mode. In Dev, it may panic on some errors.")
	flag.Parse()

	if config.IsDev {
		log.Println("Dev mode enabled")
	}

	if !config.IsDev {
		log.Println("Prod mode enabled")
		config.StaticFolder = "./staticss"
	}
}

func frontEndFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("static folder" + usr.HomeDir)
	return usr.HomeDir + "/projects/lenovo-shop/front-end"
}
