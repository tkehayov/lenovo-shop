package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

type Groups struct {
	Id   string
	Name string
}

func AddGroup(groups []Groups) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}
	keys := []*datastore.Key{}
	for _, group := range groups {

		key := datastore.NameKey("Groups", group.Id, nil)
		keys = append(keys, key)
	}

	dsClient.PutMulti(ctx, keys, groups)
}

func GetAllGroups() []Groups {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	var gr []Groups
	q := datastore.NewQuery("Groups")
	dsClient.GetAll(ctx, q, &gr)

	return gr
}
