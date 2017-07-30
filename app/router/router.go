package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/lenovo-shop/app/config"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/article", ArticleHandler)

	fmt.Print(config.StaticFolder)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticFolder)))
	return r
}

func ArticleHandler(writer http.ResponseWriter, request *http.Request) {
	value:=request.FormValue("Name")
	fmt.Println(value)

	fmt.Fprintf(writer, "love you so mucsshss")
}
