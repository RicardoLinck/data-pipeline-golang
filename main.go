package main

import (
	"log"
)

func main() {
	c := saveData(fetchData(
		prepareData(
			generateData(),
		),
	))

	for data := range c {
		log.Printf("Items saved: %+v", data)
	}
}
