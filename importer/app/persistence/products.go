package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Product struct {
	Description  string
	Model        string
	Name         string
	Kkprice      float64
	Price        float64
	PriceDds     float64
	SmallPicture string
}

type ProductId struct {
	Id string
}

func PersistMultiProducts(prodsId ...ProductId) {
	ctx := context.Background()
	var pr []*ProductId
	var keys []*datastore.Key
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}

	for _, prod := range prodsId {
		pr = append(pr, &prod)
		key := datastore.NameKey("ProductProvider", prod.Id, nil)
		keys = append(keys, key)
	}

	if _, err := dsClient.PutMulti(ctx, keys, pr); err != nil {
		log.Print(err)
	}
}

func GetAllProducts() []ProductId {
	var pr []ProductId
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}

	q := datastore.NewQuery("ProductProvider")
	dsClient.GetAll(ctx, q, &pr)

	return pr
}
