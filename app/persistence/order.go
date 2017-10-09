package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Order struct {
	Firstname      string
	Lastname       string
	Address        string
	Location       string
	Email          string
	Products       []*datastore.Key
	ProductsEntity []Product
}

func MakeOrder(order Order, ids ...int64) {
	var keys []*datastore.Key

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	//	Loop and append keys
	for _, id := range ids {
		idKey := datastore.IDKey("Products", id, nil)
		keys = append(keys, idKey)
	}
	order.Products = keys

	orderKey := datastore.IncompleteKey("Orders", nil)

	if _, err := dsClient.Put(ctx, orderKey, &order); err != nil {
		log.Fatal(err)
	}
}

func ListOrders() []Order {
	var entities []Order

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	q := datastore.NewQuery("Orders").Limit(10)
	dsClient.GetAll(ctx, q, &entities)

	for index, order := range entities {
		var product Product
		var productOrder []Product

		for _, p := range order.Products {
			dsClient.Get(ctx, p, &product)
			productOrder = append(productOrder, product)
		}

		log.Fatal("productOrder", productOrder)
		order.ProductsEntity = productOrder
		entities[index] = order
	}

	log.Print("entities", entities)

	return entities
}
