package extract_cars

import (
	"fmt"
	"mobilebg2/car"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ChannelProfiler struct {
	chan_page_with_car_links chan *goquery.Document
	chan_car_links           chan string
	chan_car_pages           chan *carPageData
	chan_cars                chan *car.Car

	blockcount_running bool

	chan_page_with_car_links_blockcount int
	chan_car_links_blockcount           int
	chan_car_pages_blockcount           int
	chan_cars_blockcount                int

	samples_taken int
}

func NewChannelProfiler(
	chan_page_with_car_links chan *goquery.Document,
	chan_car_links chan string,
	chan_car_pages chan *carPageData,
	chan_cars chan *car.Car,
) *ChannelProfiler {
	self := &ChannelProfiler{
		chan_page_with_car_links: chan_page_with_car_links,
		chan_car_links:           chan_car_links,
		chan_car_pages:           chan_car_pages,
		chan_cars:                chan_cars,

		blockcount_running: false,

		chan_page_with_car_links_blockcount: 0,
		chan_car_links_blockcount:           0,
		chan_car_pages_blockcount:           0,
		chan_cars_blockcount:                0,

		samples_taken: 0,
	}

	go self.checkBlockings()

	return self
}

func (self *ChannelProfiler) checkBlockings() {
	if self.blockcount_running {
		panic("blockcount_running")
	}

	self.blockcount_running = true

	for self.blockcount_running {
		self.samples_taken += 1

		if len(self.chan_page_with_car_links) == cap(self.chan_page_with_car_links) {
			self.chan_page_with_car_links_blockcount += 1
		}

		if len(self.chan_car_links) == cap(self.chan_car_links) {
			self.chan_car_links_blockcount += 1
		}

		if len(self.chan_car_pages) == cap(self.chan_car_pages) {
			self.chan_car_pages_blockcount += 1
		}

		if len(self.chan_cars) == cap(self.chan_cars) {
			self.chan_cars_blockcount += 1
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (self *ChannelProfiler) ShowResults() {
	self.blockcount_running = false

	time.Sleep(time.Millisecond * 100 * 2) // TODO: this is shit

	fmt.Printf("Channel Blockings:\n")
	fmt.Printf("chan_page_with_car_links: %3v / %3v | %v%%\n", self.chan_page_with_car_links_blockcount, self.samples_taken, 100*float32(self.chan_page_with_car_links_blockcount)/float32(self.samples_taken))
	fmt.Printf("chan_car_links          : %3v / %3v | %v%%\n", self.chan_car_links_blockcount, self.samples_taken, 100*float32(self.chan_car_links_blockcount)/float32(self.samples_taken))
	fmt.Printf("chan_car_pages          : %3v / %3v | %v%%\n", self.chan_car_pages_blockcount, self.samples_taken, 100*float32(self.chan_car_pages_blockcount)/float32(self.samples_taken))
	fmt.Printf("chan_cars               : %3v / %3v | %v%%\n", self.chan_cars_blockcount, self.samples_taken, 100*float32(self.chan_cars_blockcount)/float32(self.samples_taken))
}
