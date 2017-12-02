package model

type Vendor struct {
	Vendor []Vendors `xml:"vendor"`
}

type Vendors struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}
