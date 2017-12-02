package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/lenovo-shop/app/shared"
	"log"
)

type Filter struct {
	ScreenSizes []string
	Category    Category
	PriceFrom   float32
	PriceTo     float32
	SubCategory []string
	Limit       int
	OrderPrice  string
}

func FilterProducts(filter Filter) []Product {
	var products []Product
	var productsSeries []Product
	var seriesKeys []*datastore.Key
	var productKeys []*datastore.Key
	ctx, dsClient := shared.Connect()
	log.Print(filter.SubCategory)
	kCat := datastore.NameKey("Categories", filter.Category.Name, nil)

	//Product Series
	productsSeries, seriesKeys = filterSeries(filter, dsClient, ctx, productsSeries, seriesKeys)
	log.Print("productsSeries", productsSeries)

	for _, screenSize := range filter.ScreenSizes {
		q := datastore.NewQuery("Products")

		if filter.OrderPrice == "priceAsc" {
			q = q.Order("Price")
		} else {
			q = q.Order("-Price")

		}
		q = q.Limit(filter.Limit)

		if filter.Category.Name != "" {
			//q = q.Ancestor(kCat).Filter("Slug=", filter.Category.Slug)
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
	for _, series := range filter.SubCategory {
		//TODO in factory method
		querySeries := datastore.NewQuery("Products")
		if filter.OrderPrice == "priceAsc" {
			querySeries = querySeries.Order("Price")
		} else {
			querySeries = querySeries.Order("-Price")

		}
		querySeries = querySeries.Limit(filter.Limit)

		if len(series) != 0 {
			cat := datastore.NameKey("Categories", filter.Category.Name, nil)
			key := datastore.NameKey("SubCategories", series, cat)
			querySeries = querySeries.Ancestor(key)
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
