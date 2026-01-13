package extract_cars

import (
	"mobilebg2/car"
	"mobilebg2/define"
	"mobilebg2/net"
	"mobilebg2/profiler"

	"github.com/PuerkitoBio/goquery"
)

func Main(
	channelProfiler *profiler.ChannelProfiler,
	chan_net_req chan *net.ReqData,
) (
	chan_cars_ chan *car.Car,
) {
	chan_page_with_car_links := make(chan *goquery.Document, define.CHAN_BUF_PAGE_WITH_CAR_LINKS)
	chan_car_links := make(chan string, define.CHAN_BUF_CAR_LINKS)
	chan_car_pages := make(chan *carPageData, define.CHAN_BUF_CAR_PAGES)
	chan_cars := make(chan *car.Car, define.CHAN_BUF_CAR)

	channelProfiler.AddChannels(
		profiler.NewChannelData("chan_page_with_car_links", func() int { return len(chan_page_with_car_links) }, cap(chan_page_with_car_links)),
		profiler.NewChannelData("chan_car_links", func() int { return len(chan_car_links) }, cap(chan_car_links)),
		profiler.NewChannelData("chan_car_pages", func() int { return len(chan_car_pages) }, cap(chan_car_pages)),
		profiler.NewChannelData("chan_cars", func() int { return len(chan_cars) }, cap(chan_cars)),
	)

	go extractSearchPages(chan_net_req, chan_page_with_car_links)
	go extractCarLinks(chan_page_with_car_links, chan_car_links)
	go downloadCarPages(chan_net_req, chan_car_links, chan_car_pages)
	go extractCars(chan_car_pages, chan_cars)

	return chan_cars
}
