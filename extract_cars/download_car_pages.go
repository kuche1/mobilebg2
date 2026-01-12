package extract_cars

import (
	"log"
	"mobilebg2/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type carPageData struct {
	link string
	doc  *goquery.Document
}

func downloadCarPages(chan_net_req chan *net.ReqData, chan_car_links chan string, chan_car_pages chan *carPageData) {
	defer close(chan_car_pages)

	for link := range chan_car_links {
		// fmt.Printf("car page link: %v\n", link)

		page_bytes := <-net.Req(chan_net_req, link)
		page_text := string(page_bytes)

		reader0 := strings.NewReader(page_text)

		reader1, err := charset.NewReaderLabel("windows-1251", reader0)
		if err != nil {
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(reader1)
		if err != nil {
			log.Fatal(err)
		}

		chan_car_pages <- &carPageData{
			link,
			doc,
		}
	}
}
