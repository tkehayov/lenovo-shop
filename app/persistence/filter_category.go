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
	Series      []string
}

func FilterProducts(filter Filter) []Product {
	var products []Product
	var productsSeries []Product
	var seriesKeys []*datastore.Key
	var productKeys []*datastore.Key

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))

	if err != nil {
		log.Print(err)
	}

	kCat := datastore.NameKey("Categories", filter.Category, nil)

	//Product Series
	productsSeries, seriesKeys = filterSeries(filter, dsClient, ctx, productsSeries, seriesKeys)

	for _, screenSize := range filter.ScreenSizes {
		q := datastore.NewQuery("Products")

		if filter.Category != "" {
			q = q.Ancestor(kCat)
		}

		//Product screenSize
		if len(screenSize) != 0 {
			q = q.Filter("ScreenSize=", screenSize)
		}

		//Product Price
		if filter.PriceTo != 0 {
			q = q.Filter("Price>=", filter.PriceFrom).Filter("Price<=", filter.PriceTo)
		}

		keys, errf := dsClient.GetAll(ctx, q, &products)
		productKeys = append(productKeys, keys...)

		if errf != nil {
			log.Print(errf)
		}

	}

	for index, key := range productKeys {
		products[index].Id = key.ID
	}

	products = normalizeProduct(products, productsSeries)

	return products
}
func filterSeries(filter Filter, dsClient *datastore.Client, ctx context.Context, productsSeries []Product, seriesKeys []*datastore.Key) ([]Product, []*datastore.Key) {
	for _, series := range filter.Series {
		querySeries := datastore.NewQuery("Products")

		if len(series) != 0 {
			querySeries = querySeries.Filter("Series=", series)
		}

		keys, errf := dsClient.GetAll(ctx, querySeries, &productsSeries)
		seriesKeys = append(seriesKeys, keys...)

		if errf != nil {
			log.Print(errf)
		}
	}
	for index, _ := range productsSeries {
		productsSeries[index].Id = seriesKeys[index].ID
	}
	return productsSeries, seriesKeys
}

func normalizeProduct(products []Product, series []Product) []Product {
	var results []Product

	for _, ser := range series {
		for _, prod := range products {
			if prod.Id == ser.Id {
				results = append(results, prod)
			}
		}
	}

	return results
}
