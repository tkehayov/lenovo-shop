package router

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/config"
	"net/http"
	"github.com/lenovo-shop/app/model"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/article", ArticleHandler)
	r.HandleFunc("/read", ReadHandler)

	fmt.Print(config.StaticFolder)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticFolder)))
	return r
}

func ReadHandler(writer http.ResponseWriter, request *http.Request) {
	model.ReadCategories("./export.json")
}

func ArticleHandler(writer http.ResponseWriter, request *http.Request) {
	value := request.FormValue("Name")
	fmt.Println(value)

	fmt.Fprintf(writer, "love you so mucsshss")
}
