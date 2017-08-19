package cart

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type Cookie struct {
	ID          string
	ProductName string
	Quantity    int
}

func Add(w http.ResponseWriter, c Cookie) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{Name: "cart", Value: base64.StdEncoding.EncodeToString(b)}

	http.SetCookie(w, &cookie)
}

func Get(req *http.Request) (*Cookie, error) {
	cookie, cerr := req.Cookie("cart")

	if cerr != nil {
		return nil, cerr
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, err
	}

	var c *Cookie
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return c, nil
}
