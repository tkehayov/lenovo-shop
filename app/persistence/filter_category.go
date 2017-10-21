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

	for _, screenSize := range filter.ScreenSizes {
		q := datastore.NewQuery("Products").Filter("ScreenSize=", screenSize)
		_, errf := dsClient.GetAll(ctx, q, &products)

		if errf != nil {
			log.Print(errf)
		}

	}

	return products
}
