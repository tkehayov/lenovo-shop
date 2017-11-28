package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/lenovo-shop/app/shared"
	"log"
)

type Category struct {
	Name string
}

func AddCategory(category Category) {
	ctx := context.Background()
	ctx, dsClient := shared.Connect()

	key := datastore.NameKey("Categories", category.Name, nil)

	if _, err := dsClient.Put(ctx, key, &category); err != nil {
		log.Print(err)
	}

}

func GetAllCategories() []Category {
	ctx, dsClient := shared.Connect()

	var cat []Category
	q := datastore.NewQuery("Categories")
	dsClient.GetAll(ctx, q, &cat)

	return cat
}
