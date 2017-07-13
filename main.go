package main

import (
	"log"
	"flag"
	"fmt"
	"net/http"
	"github.com/lenovo-shop/app/router"
)

var port string
var isDebug bool

func main() {
	http.Handle("/", router.GetRouter())
	http.ListenAndServe(port, nil)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to home")
}

func init() {
	log.Println("Starting...")

	flag.StringVar(&port, "port", ":8080", "Specify the port to listen to e.g. :8080")
	flag.BoolVar(&isDebug, "isDebug", true, "Set to true to run the app in debug mode. In debug, it may panic on some errors.")
	flag.Parse()

	if isDebug {
		log.Println("DEBUG mode enabled")
	}

}
