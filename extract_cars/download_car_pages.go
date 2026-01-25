package extract_cars

import (
	"log"
	"mobilebg2/define"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"golang.org/x/net/html/charset"
)

type carPageData struct {
	link string
	doc  *goquery.Document
}

func downloadCarPages(net *gonet.Net, chan_car_links chan string, chan_car_pages chan<- *carPageData) {
	defer close(chan_car_pages)

	var wg sync.WaitGroup

	for range define.THREADS_DOWNLOAD_CAR_PAGES {
		wg.Go(func() {
			downloadCarPagesThread(net, chan_car_links, chan_car_pages)
		})
	}

	wg.Wait()
}

func downloadCarPagesThread(net *gonet.Net, chan_car_links chan string, chan_car_pages chan<- *carPageData) {
	for link := range chan_car_links {
		// fmt.Printf("downloadCarPages: got car page link: %v\n", link)

		page_bytes := net.Req(link)
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

		// fmt.Printf("downloadCarPages: done\n")
	}
}
