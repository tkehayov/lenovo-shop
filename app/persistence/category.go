package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Category struct {
	Name string
}

func AddCategory(category Category) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	key := datastore.NameKey("Categories", category.Name, nil)

	if _, err := dsClient.Put(ctx, key, &category); err != nil {
		log.Fatal(err)
	}

}
