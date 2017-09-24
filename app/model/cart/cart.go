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
	Id       int64
	Quantity int
}

func Add(w http.ResponseWriter, req *http.Request, c CartCookie, cc *[]CartCookie) {
	//retrieve shopping cart
	var hasQuantity bool

	cookies, errCookies := Get(req)

	//increase quantity to an already exists product
	for i := range cookies {

		if cookies[i].Id == c.Id {
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

	encodedCookie := marshalCookie(cookies)

	http.SetCookie(w, &encodedCookie)
}

func Get(req *http.Request) ([]CartCookie, error) {
	cookie, cerr := req.Cookie("shoppingcart")

	if cerr != nil {
		return nil, cerr
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	var ca []CartCookie

	if err := json.Unmarshal(data, &ca); err != nil {
		return nil, err
	}

	return ca, nil
}

func Delete(w http.ResponseWriter, req *http.Request) {
	cookie, cerr := req.Cookie("shoppingcart")
	var c []CartCookie

	vars := mux.Vars(req)
	idVar := vars["id"]
	id, errParsing := strconv.ParseInt(idVar,10,64)
	if errParsing != nil {
		log.Print(errParsing)
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
		if v.Id == id {
			c = append(c[:i], c[i+1:]...)
		}
	}

	encodedCookie := marshalCookie(c)

	http.SetCookie(w, &encodedCookie)

}

func marshalCookie(cookies []CartCookie) http.Cookie {
	b, err := json.Marshal(cookies)
	if err != nil {
		log.Fatal(err)
	}
	encodedCookie := http.Cookie{Name: "shoppingcart", Path: "/", Value: base64.StdEncoding.EncodeToString(b)}
	return encodedCookie
}
