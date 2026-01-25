package extract_cars

import (
	"fmt"
	"log"
	"mobilebg2/config"
	"mobilebg2/define"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
)

func extractSearchPages(net *gonet.Net, chan_page_with_car_links chan<- *goquery.Document) {
	defer close(chan_page_with_car_links)

	var wg sync.WaitGroup

	threadsSpawned := 0
	chanThreadDone := make(chan struct{})

	price_max := config.PRICE_MAX

	for {
		if price_max < config.PRICE_MIN {
			break
		}

		price_min := max(price_max-define.PRICE_STEP, config.PRICE_MIN)

		// this print sucks and is misleading
		// fmt.Printf(
		// 	"Extract Data For Price Range: [%v] / %v / %v / [%v]\n",
		// 	config.PRICE_MIN,
		// 	price_min,
		// 	price_max,
		// 	config.PRICE_MAX,
		// )

		// otherwise we are going to capture references ot the variables, and by the time
		// the thread has started the values will have changed
		anon_price_min := price_min
		anon_price_max := price_max

		wg.Go(func() {
			extractSearchPagesWithinPriceRange(net, chan_page_with_car_links, anon_price_min, anon_price_max, chanThreadDone)
		})
		threadsSpawned += 1

		price_max = price_min - 1
	}

	// fmt.Printf("Extract Data For Price Range: Done\n")

	for {
		// fmt.Printf("Loading: %v tasks left\n", threadsSpawned)

		if threadsSpawned == 0 {
			break
		}

		<-chanThreadDone

		threadsSpawned -= 1
	}

	wg.Wait()

	// fmt.Printf("Loading: Done\n")
}

func extractSearchPagesWithinPriceRange(
	net *gonet.Net,
	chan_page_with_car_links chan<- *goquery.Document,
	price_min int,
	price_max int,
	chanThreadDone chan<- struct{},
) {
	defer func() { chanThreadDone <- struct{}{} }()

	var page_num int

	for page_num = 1; page_num <= define.SEARCH_MAX_PAGE; page_num++ {

		url := fmt.Sprintf(define.SEARCH_URL, page_num, price_min, price_max)
		// fmt.Printf("Car Pages: Url: %v | price_min=%v price_max=%v\n", url, price_min, price_max)

		raw_resp_bytes := net.Req(url)
		raw_resp_text := string(raw_resp_bytes)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw_resp_text))
		if err != nil {
			log.Fatal(err)
		}

		// find the "Няма намерени обяви!" message
		if doc.Find("div.width980px.pageMessageAlert").Length() > 0 {
			break
		}

		chan_page_with_car_links <- doc
	}

	if page_num >= define.SEARCH_MAX_PAGE {
		log.Printf("The very last search page was reached, it is possible that some cars were omitted")
	}

	// fmt.Printf("Car Pages: Done\n")
}
