package extract_cars

import (
	"fmt"
	"mobilebg2/configuration"
	"mobilebg2/define"
	"slices"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func extractCarLinks(chan_page_with_car_links <-chan *goquery.Document, chan_car_links chan<- string, config *configuration.Config) {
	defer close(chan_car_links)

	var wg sync.WaitGroup

	for range define.THREADS_EXTRACT_CAR_LINKS {
		wg.Go(func() {
			extractCarLinksThr(chan_page_with_car_links, chan_car_links, config)
		})
	}

	wg.Wait()
}

func extractCarLinksThr(chan_page_with_car_links <-chan *goquery.Document, chan_car_links chan<- string, config *configuration.Config) {
	for doc := range chan_page_with_car_links {
		// fmt.Printf("Extract Car Links: %v\n", doc)

		doc.Find("a.title.saveSlink").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if !exists {
				fmt.Printf("Could not find href\n")
				return
			}

			// fmt.Printf("href1: %v\n", href)

			href = define.CAR_LINK_PREFIX + href

			// fmt.Printf("href2: %v\n", href)

			if slices.Contains(config.LinkBlacklist, href) {
				return
			}

			chan_car_links <- href
		})

		// fmt.Printf("Extract Car Links: Done\n")
	}

	// fmt.Printf("Extract Car Links: Done\n")
}
