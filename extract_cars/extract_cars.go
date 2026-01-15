package extract_cars

import (
	"fmt"
	"log"
	"mobilebg2/car"
	"mobilebg2/config"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractCars(chan_car_pages chan *carPageData, chan_cars chan *car.Car) {
	defer close(chan_cars)

	// fmt.Printf("extractCars: Begin\n")

	for page_data := range chan_car_pages {
		// fmt.Printf("extractCars: Got some data\n")

		elem_info := page_data.doc.Find("div.contactsBox").First()
		elem_params := page_data.doc.Find("div.mainCarParams").First()

		title, blacklisted := findTitle(elem_info)
		if blacklisted {
			// fmt.Printf("Blacklisted car: %v\n", title)
			continue
		}

		price, invalid := findPrice(elem_info, page_data.link)
		if invalid {
			continue
		}

		engineType, blacklisted := findEngineType(elem_params)
		if blacklisted {
			continue
		}

		horsepower, blacklisted := findHorsepower(elem_params)
		if blacklisted {
			continue
		}

		// TODO: missing reseller

		// TODO: missing gearbox

		// TODO: missing mialage

		// TODO: missing date produced

		// TODO: missing description

		chan_cars <- car.NewCar(
			page_data.link,
			title,
			price,
			engineType,
			horsepower,
		)

		// fmt.Printf("extract_cars: processed data\n")
	}

	// fmt.Printf("extractCars: End\n")
}

func findTitle(elem_info *goquery.Selection) (value string, blacklisted bool) {
	elem_title := elem_info.Find("div.obTitle").First()

	title := strings.TrimSpace(elem_title.Text())
	parts := strings.Split(title, " Обява: ")
	title = parts[0]

	title_lower := strings.ToLower(title)

	if len(config.TITLE_PREFIX_WHITELIST) > 0 {
		found := false

		for _, whitelisted_title := range config.TITLE_PREFIX_WHITELIST {
			whitelisted_title_lower := strings.ToLower(whitelisted_title)

			if strings.HasPrefix(title_lower, whitelisted_title_lower) {
				found = true
				break
			}
		}

		if !found {
			return title, true
		}
	}

	for _, blacklisted_title := range config.TITLE_PREFIX_BLACKLIST {
		blacklisted_title_lower := strings.ToLower(blacklisted_title)

		if strings.HasPrefix(title_lower, blacklisted_title_lower) {
			return title, true
		}
	}

	return title, false
}

func findPrice(elem_info *goquery.Selection, url string) (_price float64, _invalid bool) {
	const eur = "€"

	elem := elem_info.Find("div.Price").First()

	price := strings.TrimSpace(elem.Text())

	if !strings.Contains(price, eur) {
		log.Printf("Price in euros not found: %v", url)
		return 0, true
	}

	parts := strings.Split(price, eur)

	price = strings.TrimSpace(parts[0])
	price = strings.ReplaceAll(price, " ", "")

	value, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Fatal("URL `", url, "`:", err)
	}

	return value, false
}

func findEngineType(elem_params *goquery.Selection) (_engineType string, _blacklisted bool) {
	elem := elem_params.Find("div.item.dvigatel").First()
	// NOTE: original python code: soup.find("div", class_="item dvigatel")

	elem = elem.Find("div.mpInfo")

	engineType := elem.Text()

	if slices.Contains(config.ENGINE_TYPE_BLACKLIST, engineType) {
		return "", true
	}

	return engineType, false
}

func findHorsepower(elem_params *goquery.Selection) (_value int64, _blacklisted bool) {
	elem := elem_params.Find("div.item.moshtnost").First()
	// NOTE: original python code: soup.find("div", class_="item moshtnost")

	if elem.Length() == 0 {
		if config.HORSEPOWER_MISSING_OK {
			return -1, false
		}
		return -1, true
	}

	elem = elem.Find("div.mpInfo")

	valueAsStr := elem.Text()

	suffix := " к.с."
	if !strings.HasSuffix(valueAsStr, suffix) {
		panic("The site must have changed")
	}
	valueAsStr = strings.TrimSuffix(valueAsStr, suffix)

	valueAsInt, err := strconv.ParseInt(valueAsStr, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("The site must have changed: %v", err))
	}

	if valueAsInt < config.HORSEPOWER_MIN {
		return -1, true
	}

	return valueAsInt, false
}
