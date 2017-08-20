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

func Add(w http.ResponseWriter, c CartCookie) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{Name: "cart", Value: base64.StdEncoding.EncodeToString(b)}

	http.SetCookie(w, &cookie)
}

func Get(req *http.Request) (*CartCookie, error) {
	cookie, cerr := req.Cookie("cart")

	if cerr != nil {
		return nil, cerr
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	var c *CartCookie
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return c, nil
}
