package persistence


type Product struct {
	ID    int
	Price float32
	Name  string
}

func Persist(pr Product) {
	//log.Print(pr)
	//
	//db, err := sql.Open("mysql", config.DbUri)
	//if err != nil {
	//	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	//}
	//defer db.Close()
	//
	//// Open doesn't open a connection. Validate DSN data:
	//err = db.Ping()
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//
	//stmtIns, err := db.Prepare("INSERT INTO products VALUES( ?, ? )")
	//if err != nil {
	//	log.Print(err)
	//}
	//defer stmtIns.Close()
	//
	//stmtIns.Exec(pr.Name, 2.31)
	//
	//fmt.Println("success")
}

func Get(ids ...int) []Product {
	var product []Product

	//db, err := sql.Open("mysql", config.DbUri)
	//if err != nil {
	//
	//	log.Fatal(err.Error())
	//}
	//
	//args := []interface{}{}
	//for _, value := range ids {
	//	args = append(args, value)
	//}
	//
	//selDB, err := db.Query("SELECT id,name,price FROM products WHERE id in (?"+strings.Repeat(",?", len(ids)-1)+")", args...)
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//index := 0
	//for selDB.Next() {
	//	product = append(product, Product{})
	//	err = selDB.Scan(&product[index].ID, &product[index].Name, &product[index].Price)
	//	index++
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//defer db.Close()

	return product
}
