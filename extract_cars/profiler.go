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

	chan_page_with_car_links_fullcount int
	chan_car_links_fullcount           int
	chan_car_pages_fullcount           int
	chan_cars_fullcount                int

	chan_page_with_car_links_emptycount int
	chan_car_links_emptycount           int
	chan_car_pages_emptycount           int
	chan_cars_emptycount                int

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

		chan_page_with_car_links_fullcount: 0,
		chan_car_links_fullcount:           0,
		chan_car_pages_fullcount:           0,
		chan_cars_fullcount:                0,

		chan_page_with_car_links_emptycount: 0,
		chan_car_links_emptycount:           0,
		chan_car_pages_emptycount:           0,
		chan_cars_emptycount:                0,

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

		// TODO: need to take the lens here in variables

		///// full

		if len(self.chan_page_with_car_links) == cap(self.chan_page_with_car_links) {
			self.chan_page_with_car_links_fullcount += 1
		}

		if len(self.chan_car_links) == cap(self.chan_car_links) {
			self.chan_car_links_fullcount += 1
		}

		if len(self.chan_car_pages) == cap(self.chan_car_pages) {
			self.chan_car_pages_fullcount += 1
		}

		if len(self.chan_cars) == cap(self.chan_cars) {
			self.chan_cars_fullcount += 1
		}

		///// empty

		if len(self.chan_page_with_car_links) == 0 {
			self.chan_page_with_car_links_emptycount += 1
		}

		if len(self.chan_car_links) == 0 {
			self.chan_car_links_emptycount += 1
		}

		if len(self.chan_car_pages) == 0 {
			self.chan_car_pages_emptycount += 1
		}

		if len(self.chan_cars) == 0 {
			self.chan_cars_emptycount += 1
		}

		/////

		time.Sleep(time.Millisecond * 100)
	}
}

func (self *ChannelProfiler) ShowResults() {
	self.blockcount_running = false

	time.Sleep(time.Millisecond * 100 * 3) // TODO: this is shit

	fmt.Printf("\n")
	fmt.Printf("Channel Full Counts:\n")
	fmt.Printf("chan_page_with_car_links: %3v / %3v | %6.2f%%\n", self.chan_page_with_car_links_fullcount, self.samples_taken, 100*float32(self.chan_page_with_car_links_fullcount)/float32(self.samples_taken))
	fmt.Printf("chan_car_links          : %3v / %3v | %6.2f%%\n", self.chan_car_links_fullcount, self.samples_taken, 100*float32(self.chan_car_links_fullcount)/float32(self.samples_taken))
	fmt.Printf("chan_car_pages          : %3v / %3v | %6.2f%%\n", self.chan_car_pages_fullcount, self.samples_taken, 100*float32(self.chan_car_pages_fullcount)/float32(self.samples_taken))
	fmt.Printf("chan_cars               : %3v / %3v | %6.2f%%\n", self.chan_cars_fullcount, self.samples_taken, 100*float32(self.chan_cars_fullcount)/float32(self.samples_taken))
	fmt.Printf("\n")
	fmt.Printf("Channel Empty Counts:\n")
	fmt.Printf("chan_page_with_car_links: %3v / %3v | %6.2f%%\n", self.chan_page_with_car_links_emptycount, self.samples_taken, 100*float32(self.chan_page_with_car_links_emptycount)/float32(self.samples_taken))
	fmt.Printf("chan_car_links          : %3v / %3v | %6.2f%%\n", self.chan_car_links_emptycount, self.samples_taken, 100*float32(self.chan_car_links_emptycount)/float32(self.samples_taken))
	fmt.Printf("chan_car_pages          : %3v / %3v | %6.2f%%\n", self.chan_car_pages_emptycount, self.samples_taken, 100*float32(self.chan_car_pages_emptycount)/float32(self.samples_taken))
	fmt.Printf("chan_cars               : %3v / %3v | %6.2f%%\n", self.chan_cars_emptycount, self.samples_taken, 100*float32(self.chan_cars_emptycount)/float32(self.samples_taken))
}
