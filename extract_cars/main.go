package extract_cars

import (
	"mobilebg2/car"
	"mobilebg2/config"
	"mobilebg2/net"

	"github.com/PuerkitoBio/goquery"
)

func Main(chan_net_req chan *net.ReqData) (chan_cars chan *car.Car) {
	chan_page_with_car_links := make(chan *goquery.Document, config.CHAN_BUF_PAGE_WITH_CAR_LINKS)
	chan_car_links := make(chan string, config.CHAN_BUF_CAR_LINKS)
	chan_car_pages := make(chan *carPageData, config.CHAN_BUF_CAR_PAGES)
	chan_cars = make(chan *car.Car, config.CHAN_BUF_CAR)

	go extractSearchPages(chan_net_req, chan_page_with_car_links)
	go extractCarLinks(chan_page_with_car_links, chan_car_links)
	go downloadCarPages(chan_net_req, chan_car_links, chan_car_pages)
	go extractCars(chan_car_pages, chan_cars)

	return
}
