package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Product struct {
	ID    int64
	Price float32
	Name  string
}

func Persist(pr Product) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	productKey := datastore.IncompleteKey("Products", nil)

	products := &Product{Price: pr.Price, Name: pr.Name}
	if _, err := dsClient.Put(ctx, productKey, products); err != nil {
		log.Fatal(err)
	}
}

func Get(keysID ...int64) []Product {
	products := make([]Product, len(keysID))
	keys := []*datastore.Key{}

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Fatal(err)
	}

	for _, id := range keysID {
		k := datastore.IDKey("Products", id, nil)
		keys = append(keys, k)
	}

	errProducts := dsClient.GetMulti(ctx, keys, products)
	if errProducts != nil {
		log.Fatal(errProducts)
	}

	return products
}

func GetAll() []Product {
	var products []Product

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Fatal(err)
	}
	q := datastore.NewQuery("Products")

	keys, erra := dsClient.GetAll(ctx, q, &products)

	if erra != nil {
		log.Fatal(erra)
	}

	for index, k := range keys {
		products[index].ID = k.ID
	}
	log.Fatal(products)

	return products
}
