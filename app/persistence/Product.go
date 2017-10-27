package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Product struct {
	Id           int64
	Price        float32
	Name         string
	ScreenSize   string
	ImagePreview string
	Category     string
}

func Persist(pr Product) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	par := datastore.NameKey("Category", pr.Category, nil)
	productKey := datastore.IncompleteKey("Products", par)

	products := &Product{Price: pr.Price, Name: pr.Name, ScreenSize: pr.ScreenSize, ImagePreview: pr.ImagePreview}
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
		log.Fatal("errProducts", errProducts)
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
	k := datastore.NameKey("Category", "лаптопи", nil)
	q := datastore.NewQuery("Products").Ancestor(k)

	keys, erra := dsClient.GetAll(ctx, q, &products)

	if erra != nil {
		log.Fatal(erra)
	}

	for index, k := range keys {
		products[index].Id = k.ID
		products[index].Category = k.Parent.Name
	}
	log.Print(products)

	return products
}
