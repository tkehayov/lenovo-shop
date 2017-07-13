package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/article", ArticleHandler)

	return r
}
func ArticleHandler(writer http.ResponseWriter, request *http.Request) {

}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

}
