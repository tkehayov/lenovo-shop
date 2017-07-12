package main

import (
	"fmt"
	"log"
	"flag"
)

var port int
var isDebug bool

func main() {
	fmt.Print("hello from mains")
}

func init() {
	log.Println("Starting...")

	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")
	flag.BoolVar(&isDebug, "isDebug", true, "Set to true to run the app in debug mode.  In debug, it may panic on some errors.")
	flag.Parse()

	log.Println(isDebug)
	// log our flags
	if isDebug {
		log.Println("DEBUG mode enabled")
	}

}
