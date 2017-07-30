package router

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/config"
	"net/http"
	"database/sql"
	"github.com/lenovo-shop/app/model"
)

func GetRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/article", ArticleHandler)
	r.HandleFunc("/db", DbHandler)
	r.HandleFunc("/read", ReadHandler)

	fmt.Print(config.StaticFolder)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticFolder)))
	return r
}

func ReadHandler(writer http.ResponseWriter, request *http.Request) {
	model.ReadCategories("./export.json")
}

func DbHandler(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", "root:titi89@/lenovo-shop")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO products VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	stmtIns.Exec("Lsenovo", 2.31)
	//
	fmt.Println("success")
}

func ArticleHandler(writer http.ResponseWriter, request *http.Request) {
	value := request.FormValue("Name")
	fmt.Println(value)

	fmt.Fprintf(writer, "love you so mucsshss")
}
