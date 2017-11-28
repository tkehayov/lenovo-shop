package router

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/controller"
	"github.com/lenovo-shop/app/shared"
	"net/http"
)

func GetRouter(mode shared.Mode) http.Handler {
	r := mux.NewRouter()

	var getMode = func(f http.Handler) http.HandlerFunc {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "mode", mode)

			f.ServeHTTP(w, r)
		})
	}

	// Cart
	r.HandleFunc("/cart", controller.AddCart).Methods("POST")
	r.HandleFunc("/cart", controller.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", controller.DeleteCart).Methods("DELETE")

	// Orders
	r.HandleFunc("/order", controller.Order).Methods("POST")
	r.HandleFunc("/order", controller.ListOrders).Methods("GET")

	// Products
	r.HandleFunc("/product", controller.AddProduct).Methods("POST")

	r.HandleFunc("/product/{id}", controller.GetProduct).Methods("GET")
	//TODO Not using in productions(for test purpose because of missing datastore UI)

	r.HandleFunc("/products/all", controller.GetAllProduct).Methods("GET")

	//Categories
	r.HandleFunc("/categories/{category}", getMode(http.HandlerFunc(controller.FilterProducts))).Methods("GET")
	r.HandleFunc("/categories", controller.AddCategory).Methods("POST")
	//For Testing
	r.HandleFunc("/categoriess/all", controller.GetAllCategories).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(mode.StaticPath())))

	return r
}
