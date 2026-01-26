package extract_cars

import (
	"mobilebg2/car"
	"mobilebg2/configuration"
	"mobilebg2/define"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/channelprofiler"
	"github.com/kuche1/gonet"
)

func Main(
	channelProfiler *channelprofiler.ChannelProfiler,
	net *gonet.Net,
	config *configuration.Config,
) (
	chan_cars_ chan *car.Car,
) {
	chan_page_with_car_links := make(chan *goquery.Document, define.CHAN_BUF_PAGE_WITH_CAR_LINKS)
	chan_car_links := make(chan string, define.CHAN_BUF_CAR_LINKS)
	chan_car_pages := make(chan *carPageData, define.CHAN_BUF_CAR_PAGES)
	chan_cars := make(chan *car.Car, define.CHAN_BUF_CAR)

	channelProfiler.AddChannels(
		channelprofiler.NewChannelData(
			"chan_page_with_car_links",
			func() int { return len(chan_page_with_car_links) },
			cap(chan_page_with_car_links),
		),
		channelprofiler.NewChannelData(
			"chan_car_links",
			func() int { return len(chan_car_links) },
			cap(chan_car_links),
		),
		channelprofiler.NewChannelData(
			"chan_car_pages",
			func() int { return len(chan_car_pages) },
			cap(chan_car_pages),
		),
		channelprofiler.NewChannelData(
			"chan_cars",
			func() int { return len(chan_cars) },
			cap(chan_cars),
		),
	)

	go extractSearchPages(net, chan_page_with_car_links, config)
	go extractCarLinks(chan_page_with_car_links, chan_car_links)
	go downloadCarPages(net, chan_car_links, chan_car_pages)
	go extractCars(chan_car_pages, chan_cars)

	return chan_cars
}
