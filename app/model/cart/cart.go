package cart

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type CartCookie struct {
	ID       int
	Quantity int
}

func Add(w http.ResponseWriter, req *http.Request, c CartCookie, cc *[]CartCookie) {
	//retrieve shopping cart
	var hasQuantity bool

	cookies, errCookies := Get(req)

	//increase quantity to an already exists product
	for i := range cookies {
		if cookies[i].ID == c.ID {
			cookies[i].Quantity++
			hasQuantity = true
			break
		}
	}

	if !hasQuantity || errCookies != nil {
		cookies = append(cookies, c)
	}

	*cc = cookies
	//add shopping cart
	b, err := json.Marshal(cookies)
	if err != nil {
		log.Fatal(err)
	}
	encodedCookie := http.Cookie{Name: "cart", Value: base64.StdEncoding.EncodeToString(b)}

	http.SetCookie(w, &encodedCookie)
}

func Get(req *http.Request) ([]CartCookie, error) {
	cookie, cerr := req.Cookie("cart")

	if cerr != nil {
		return nil, cerr
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	var c []CartCookie

	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return c, nil
}

func Delete(w http.ResponseWriter, req *http.Request) {
	cookie, cerr := req.Cookie("cart")
	var c []CartCookie

	vars := mux.Vars(req)
	idVar := vars["id"]
	id, errAtomic := strconv.Atoi(idVar)
	if errAtomic != nil {
		log.Print(errAtomic)
	}

	if cerr != nil {
		log.Print(cerr)
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Print(cerr)
	}

	if err := json.Unmarshal(data, &c); err != nil {
		log.Print(err)
	}

	for i, v := range c {
		if v.ID == id {
			c = append(c[:i], c[i+1:]...)
		}
	}
	//todo remove duplication
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	encodedCookie := http.Cookie{Name: "cart", Value: base64.StdEncoding.EncodeToString(b)}

	http.SetCookie(w, &encodedCookie)

}
