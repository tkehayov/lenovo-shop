package router

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/config"
	"github.com/lenovo-shop/app/controller"
	"net/http"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/article", ArticleHandler)
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticFolder)))
	return r
}

func ArticleHandler(writer http.ResponseWriter, request *http.Request) {
	value := request.FormValue("Name")
	fmt.Println(value)

}
