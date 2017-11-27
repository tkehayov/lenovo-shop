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
type SubGroups struct {
	Id   string
	Name string
}

func GetGroup(id string) Groups {
	var gr Groups
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}

	k := datastore.NameKey("CategoryProvider", id, nil)
	dsClient.Get(ctx, k, &gr)
	return gr
}

func AddGroup(groups []Groups) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}
	keys := []*datastore.Key{}
	for _, group := range groups {
		key := datastore.NameKey("GroupsProvider", group.Id, nil)
		keys = append(keys, key)
	}

	_, erra := dsClient.PutMulti(ctx, keys, groups)

	if erra != nil {
		log.Print("erra", erra)
	}
}

func GetAllGroups() []Groups {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	var gr []Groups
	q := datastore.NewQuery("GroupsProvider")
	dsClient.GetAll(ctx, q, &gr)

	return gr
}

func AddSubGroups(unmSubGroups []SubGroups) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	keys := []*datastore.Key{}

	for _, subGroups := range unmSubGroups {

		key := datastore.NameKey("SubCategoriesProvider", subGroups.Id, nil)
		keys = append(keys, key)
	}

	dsClient.PutMulti(ctx, keys, unmSubGroups)
}

func GetAllSubGroups(groups []Groups) []SubGroups {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Print(err)
	}

	var subGroups []SubGroups
	//groups

	q := datastore.NewQuery("SubGroupsProvider")
	dsClient.GetAll(ctx, q, &subGroups)

	return subGroups
}
