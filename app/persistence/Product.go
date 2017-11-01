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
	Series       string
}

func Persist(pr Product) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	cat := datastore.NameKey("Categories", pr.Category, nil)
	productKey := datastore.IncompleteKey("Products", cat)

	products := &Product{Price: pr.Price, Name: pr.Name, ScreenSize: pr.ScreenSize, ImagePreview: pr.ImagePreview, Category: pr.Category, Series: pr.Series}
	if _, err := dsClient.Put(ctx, productKey, products); err != nil {
		log.Print(err)
	}
}

func Get(keysID ...int64) []Product {
	products := make([]Product, len(keysID))
	keys := []*datastore.Key{}

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}

	for _, id := range keysID {
		k := datastore.IDKey("Products", id, nil)
		keys = append(keys, k)
	}

	errProducts := dsClient.GetMulti(ctx, keys, products)
	if errProducts != nil {
		log.Print("errProducts", errProducts)
	}

	return products
}

func GetAll() []Product {
	var products []Product

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}
	kCat := datastore.NameKey("Categories", "laptops", nil)
	q := datastore.NewQuery("Products").Ancestor(kCat)

	keys, erra := dsClient.GetAll(ctx, q, &products)

	if erra != nil {
		log.Print(erra)
	}

	for index, k := range keys {
		products[index].Id = k.ID
	}
	log.Print(products)

	return products
}
