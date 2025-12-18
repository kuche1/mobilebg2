package extract_cars

import (
	"mobilebg2/car"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractCars(chan_car_pages chan *carPageData, chan_cars chan *car.Car) {
	defer close(chan_cars)

	for page_data := range chan_car_pages {
		elem_info := page_data.doc.Find("div.contactsBox").First()

		title := findTitle(elem_info)

		chan_cars <- car.NewCar(page_data.link, title)
	}
}

func findTitle(elem_info *goquery.Selection) string {
	elem_title := elem_info.Find("div.obTitle").First()

	title := strings.TrimSpace(elem_title.Text())
	parts := strings.Split(title, " Обява: ")
	title = parts[0]

	return title
}
