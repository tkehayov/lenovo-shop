package cart

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type CartCookie struct {
	ID       string
	Quantity int
}

func Add(w http.ResponseWriter, req *http.Request, c CartCookie) {
	//retrieve shopping cart
	cookies, errCookies := Get(req)
	if errCookies != nil {
		cookies = []CartCookie{}
	}

	//add shopping cart
	cookies = append(cookies, c)

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

	//for i,v := range data{
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	//}

	return c, nil
}
