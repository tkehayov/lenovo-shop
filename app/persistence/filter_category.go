package persistence

import "log"

type Filter struct {
	ScreenSizes []string
	Category    string
}

func FilterProducts(filter Filter) {
	log.Print(filter)
}
