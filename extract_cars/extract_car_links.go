package extract_cars

import (
	"fmt"
	"mobilebg2/define"

	"github.com/PuerkitoBio/goquery"
)

func extractCarLinks(chan_page_with_car_links chan *goquery.Document, chan_car_links chan string) {
	defer close(chan_car_links)

	for doc := range chan_page_with_car_links {
		// fmt.Printf("got doc: %v\n", doc)

		doc.Find("a.title.saveSlink").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if !exists {
				fmt.Printf("Could not find href\n")
				return
			}

			// fmt.Printf("href1: %v\n", href)

			href = define.CAR_LINK_PREFIX + href

			// fmt.Printf("href2: %v\n", href)

			chan_car_links <- href
		})
	}

}
