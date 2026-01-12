package extract_cars

import (
	"mobilebg2/car"
	"mobilebg2/config"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractCars(chan_car_pages chan *carPageData, chan_cars chan *car.Car) {
	defer close(chan_cars)

	for page_data := range chan_car_pages {
		elem_info := page_data.doc.Find("div.contactsBox").First()

		title, blacklisted := findTitle(elem_info)
		if blacklisted {
			continue
		}

		chan_cars <- car.NewCar(page_data.link, title)
	}
}

func findTitle(elem_info *goquery.Selection) (value string, blacklisted bool) {
	elem_title := elem_info.Find("div.obTitle").First()

	title := strings.TrimSpace(elem_title.Text())
	parts := strings.Split(title, " Обява: ")
	title = parts[0]

	title_lower := strings.ToLower(title)

	for _, blacklisted_title := range config.BLACKLIST_TITLE_PREFIX {
		blacklisted_title_lower := strings.ToLower(blacklisted_title)

		if strings.HasPrefix(title_lower, blacklisted_title_lower) {
			return "", true
		}
	}

	return title, false
}
