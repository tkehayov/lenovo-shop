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
	PriceFrom   float32
	PriceTo     float32
}

func FilterProducts(filter Filter) []Product {
	products := []Product{}
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Fatal(err)
	}
	log.Print(filter.ScreenSizes)
	for _, screenSize := range filter.ScreenSizes {
		q := datastore.NewQuery("Products").Filter("ScreenSize=", screenSize).Filter("Price>=", filter.PriceFrom).Filter("Price<=", filter.PriceTo)
		_, errf := dsClient.GetAll(ctx, q, &products)

		if errf != nil {
			log.Print(errf)
		}

	}

	return products
}
