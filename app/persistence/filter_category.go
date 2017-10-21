package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Filter struct {
	ScreenSizes []string
	Category    string
}

func FilterProducts(filter Filter) []Product {
	products := []Product{}
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Fatal(err)
	}

	q := datastore.NewQuery("Products").Filter("ScreenSize=", filter.ScreenSizes[0])
	_, errf := dsClient.GetAll(ctx, q, &products)

	if errf != nil {
		log.Print(errf)
	}

	if len(products) == 0 {
		log.Print("No products found")
	}

	return products
}
