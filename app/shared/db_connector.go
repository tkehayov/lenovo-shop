package shared

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

func Connect() (context.Context, *datastore.Client) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	return ctx, dsClient
}
