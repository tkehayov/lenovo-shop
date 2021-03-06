package persistence

import (
	"cloud.google.com/go/datastore"
	"github.com/lenovo-shop/app/shared"
	"log"
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
	ctx, dsClient := shared.Connect()

	//	Loop and append keys
	for _, id := range ids {
		idKey := datastore.IDKey("Products", id, nil)
		keys = append(keys, idKey)
	}
	order.Products = keys

	orderKey := datastore.IncompleteKey("Orders", nil)

	if _, err := dsClient.Put(ctx, orderKey, &order); err != nil {
		log.Print(err)
	}
}

func ListOrders() []Order {
	var entities []Order
	ctx, dsClient := shared.Connect()

	q := datastore.NewQuery("Orders").Limit(10)
	dsClient.GetAll(ctx, q, &entities)

	for index, order := range entities {
		var product Product
		var productOrder []Product

		for _, p := range order.Products {
			dsClient.Get(ctx, p, &product)
			productOrder = append(productOrder, product)
		}

		log.Print("productOrder", productOrder)
		order.ProductsEntity = productOrder
		entities[index] = order
	}

	log.Print("entities", entities)

	return entities
}
