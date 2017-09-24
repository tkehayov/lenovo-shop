package persistence

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"log"
)

type Product struct {
	ID    int
	Price float32
	Name  string
}

func Persist(pr Product) {
	ctx := appengine.BackgroundContext()

	k := datastore.NewKey(ctx, "Products", "", 0, nil)

	log.Print(pr)

	if _, err := datastore.Put(ctx, k, &pr); err != nil {
		log.Print(err)
	}
}

func Get(keysID ...int64) []Product {
	products := make([]Product, len(keysID))

	ctx := appengine.BackgroundContext()
	keys := []*datastore.Key{}

	for _, keyID := range keysID {
		ka := datastore.NewKey(ctx, "Products", "", keyID, nil)
		keys = append(keys, ka)
	}
	if err := datastore.GetMulti(ctx, keys, products); err != nil {
		log.Print(err)
	}

	return products
}
