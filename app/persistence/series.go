package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Series struct {
	Name     string
	Category Category
}

func AddSeries(series Series) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	parent := datastore.NameKey("Categories", series.Category.Name, nil)
	key := datastore.NameKey("Series", series.Name, parent)

	if _, err := dsClient.Put(ctx, key, &series); err != nil {
		log.Fatal(err)
	}

}
