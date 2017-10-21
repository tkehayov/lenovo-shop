package controller

import (
	"github.com/gorilla/mux"
	"github.com/lenovo-shop/app/persistence"
	"net/http"
	"strings"
)

func FilterProducts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]

	params := req.URL.Query()
	screenSizesParam := params.Get("screenSizes")
	screenSizes := strings.Split(screenSizesParam, ",")

	filter := persistence.Filter{ScreenSizes: screenSizes, Category: category}

	persistence.FilterProducts(filter)
}
