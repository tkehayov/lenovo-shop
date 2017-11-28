package persistence

import (
	"cloud.google.com/go/datastore"
	"github.com/lenovo-shop/app/shared"
	"log"
)

type Product struct {
	Id           int64
	Price        float64
	PriceVat     float64
	Name         string
	ScreenSize   string
	ImagePreview string
	Category     string
	SubCategory  string
	Warranty     int16
	Model        string
	//Characteristics Characteristics `xml:"characteristics"`
}

func Persist(pr Product) {
	ctx, dsClient := shared.Connect()

	cat := datastore.NameKey("Categories", pr.Category, nil)
	subCat := datastore.NameKey("SubCategories", pr.SubCategory, cat)
	productKey := datastore.IncompleteKey("Products", subCat)

	products := &Product{
		Price:        pr.Price,
		Name:         pr.Name,
		ScreenSize:   pr.ScreenSize,
		ImagePreview: pr.ImagePreview,
		Category:     pr.Category,
		SubCategory:  pr.SubCategory,
	}

	if _, err := dsClient.Put(ctx, productKey, products); err != nil {
		log.Print(err)
	}
}

func PersistMulti(pr []Product) {
	var keys []*datastore.Key

	ctx, dsClient := shared.Connect()

	for _, p := range pr {
		log.Print("Category", p.Category)
		log.Print("SubCategories", p.SubCategory)
		cat := datastore.NameKey("Categories", p.Category, nil)
		subCat := datastore.NameKey("SubCategories", p.SubCategory, cat)
		productKey := datastore.IncompleteKey("Products", subCat)

		keys = append(keys, productKey)
	}

	if _, err := dsClient.PutMulti(ctx, keys, pr); err != nil {
		log.Print(err)
	}
}

func GetMulti(keysID ...int64) []Product {
	products := make([]Product, len(keysID))
	keys := []*datastore.Key{}
	ctx, dsClient := shared.Connect()

	for _, id := range keysID {
		k := datastore.IDKey("Products", id, nil)
		keys = append(keys, k)
	}

	errProducts := dsClient.GetMulti(ctx, keys, products)
	if errProducts != nil {
		log.Print("errProducts", errProducts)
	}

	return products
}

func Get(keyID int64) Product {
	var product Product

	ctx, dsClient := shared.Connect()

	kCat := datastore.NameKey("Categories", "laptops", nil)
	k := datastore.IDKey("Products", keyID, kCat)

	errProducts := dsClient.Get(ctx, k, &product)
	product.Id = k.ID

	if errProducts != nil {
		log.Print("errProduct", errProducts)
	}

	return product
}

func GetAll() []Product {
	var products []Product

	ctx, dsClient := shared.Connect()
	kCat := datastore.NameKey("Categories", "HP компютри - Pavilion", nil)
	//subCat := datastore.NameKey("SubCategories", "x1-carbon", nil)
	q := datastore.NewQuery("Products").Ancestor(kCat)

	keys, erra := dsClient.GetAll(ctx, q, &products)

	if erra != nil {
		log.Print(erra)
	}

	for index, k := range keys {
		products[index].Id = k.ID
	}
	log.Print(products)

	return products
}
