package extract_car_links

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func extractCarData(chan_docs chan *goquery.Document) {
	for doc := range chan_docs {
		fmt.Printf("got doc: %v\n", doc)
	}
}
