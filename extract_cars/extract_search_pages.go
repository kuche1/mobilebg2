package extract_cars

import (
	"fmt"
	"log"
	"mobilebg2/config"
	"mobilebg2/define"
	"mobilebg2/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractSearchPages(chan_net_req chan *net.ReqData, chan_page_with_car_links chan<- *goquery.Document) {
	defer close(chan_page_with_car_links)

	price_max := config.PRICE_MAX

	for {
		if price_max < config.PRICE_MIN {
			break
		}

		price_min := max(price_max-define.PRICE_STEP, config.PRICE_MIN)

		fmt.Printf(
			"Extract Data For Price Range: [%v] / %v / %v / [%v]\n",
			config.PRICE_MIN,
			price_min,
			price_max,
			config.PRICE_MAX,
		)

		// TODO: this needs to be parallised
		extractSearchPagesWithinPriceRange(chan_net_req, chan_page_with_car_links, price_min, price_max)

		price_max = price_min - 1
	}

	fmt.Printf("Extract Data For Price Range: Done\n")
}

func extractSearchPagesWithinPriceRange(
	chan_net_req chan *net.ReqData,
	chan_page_with_car_links chan<- *goquery.Document,
	price_min int,
	price_max int,
) {
	var page_num int

	for page_num = 1; page_num <= define.SEARCH_MAX_PAGE; page_num++ {

		url := fmt.Sprintf(define.SEARCH_URL, page_num, price_min, price_max)
		// fmt.Printf("Car Pages: Url: %v\n", url)

		raw_resp_bytes := <-net.Req(chan_net_req, url)
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
