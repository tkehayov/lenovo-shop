package main

import (
	"log"
	"flag"
	"net/http"
	"fmt"
)

var port string
var isDebug bool

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(port, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "some tsve")
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
