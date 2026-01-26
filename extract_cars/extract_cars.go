package extract_cars

import (
	"fmt"
	"log"
	"mobilebg2/car"
	"mobilebg2/configuration"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractCars(chan_car_pages chan *carPageData, chan_cars chan *car.Car, config *configuration.Config) {
	defer close(chan_cars)

	// fmt.Printf("extractCars: Begin\n")

	for page_data := range chan_car_pages {
		// fmt.Printf("extractCars: Got some data\n")

		elem_info := page_data.doc.Find("div.contactsBox").First()
		elem_params := page_data.doc.Find("div.mainCarParams").First()

		title, blacklisted := findTitle(elem_info, config)
		if blacklisted {
			// fmt.Printf("Blacklisted car: %v\n", title)
			continue
		}

		price, invalid := findPrice(elem_info, page_data.link)
		if invalid {
			continue
		}

		engineType, blacklisted := findEngineType(elem_params, config)
		if blacklisted {
			continue
		}

		horsepower, blacklisted := findHorsepower(elem_params, config)
		if blacklisted {
			continue
		}

		yearProduced, blacklisted := findYearProduced(elem_params, config)
		if blacklisted {
			continue
		}

		mialage, blacklisted := findMialage(elem_params, config)
		if blacklisted {
			continue
		}

		gearbox, blacklisted := findGearbox(elem_params, config)
		if blacklisted {
			continue
		}

		// TODO: missing reseller

		// TODO: missing description

		chan_cars <- car.NewCar(
			page_data.link,
			title,
			price,
			engineType,
			horsepower,
			yearProduced,
			mialage,
			gearbox,
		)

		// fmt.Printf("extract_cars: processed data\n")
	}

	// fmt.Printf("extractCars: End\n")
}

func findTitle(elem_info *goquery.Selection, config *configuration.Config) (value string, blacklisted bool) {
	elem_title := elem_info.Find("div.obTitle").First()

	title := strings.TrimSpace(elem_title.Text())
	parts := strings.Split(title, " Обява: ")
	title = parts[0]

	title_lower := strings.ToLower(title)

	if len(config.BrandWhitelist) > 0 {
		found := false

		for _, whitelisted_title := range config.BrandWhitelist {
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

	for _, blacklisted_title := range config.BrandBlacklist {
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

func findEngineType(elem_params *goquery.Selection, config *configuration.Config) (_engineType string, _blacklisted bool) {
	elem := elem_params.Find("div.item.dvigatel").First()
	// NOTE: original python code: soup.find("div", class_="item dvigatel")

	elem = elem.Find("div.mpInfo")

	engineType := elem.Text()

	if slices.Contains(config.EngineTypeBlacklist, engineType) {
		return "", true
	}

	return engineType, false
}

func findHorsepower(elem_params *goquery.Selection, config *configuration.Config) (_value int64, _blacklisted bool) {
	elem := elem_params.Find("div.item.moshtnost").First()
	// NOTE: original python code: soup.find("div", class_="item moshtnost")

	if elem.Length() == 0 {
		if config.HorsepowerMissingOk {
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

	if valueAsInt < config.HorsepowerMin {
		return -1, true
	}

	return valueAsInt, false
}

func findYearProduced(elem_params *goquery.Selection, config *configuration.Config) (_year int16, _blacklisted bool) {
	elem := elem_params.Find("div.item.proizvodstvo").First()
	// NOTE: original python code: elem_params.find("div", class_="item proizvodstvo")

	if elem.Length() == 0 {
		if config.YearProducedMissingOk {
			return -1, false
		}
		return -1, true
	}

	elem = elem.Find("div.mpInfo")

	monthAndYear := elem.Text()

	parts := strings.Split(monthAndYear, " ")
	if len(parts) != 2 {
		panic("Unexpected month and year format")
	}

	yearAsStr := parts[1]

	yearAsInt64, err := strconv.ParseInt(yearAsStr, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("Not a valid year: %v", err))
	}
	yearAsInt := int16(yearAsInt64)

	if yearAsInt < config.YearProducedMin {
		return -1, true
	}

	if yearAsInt > config.YearProducedMax {
		return -1, true
	}

	if len(config.YearProducedWhitelist) > 0 {
		if !slices.Contains(config.YearProducedWhitelist, yearAsInt) {
			return -1, true
		}
	}

	return yearAsInt, false
}

func findMialage(elem_params *goquery.Selection, config *configuration.Config) (_mialage int64, _blacklisted bool) {
	elem := elem_params.Find("div.item.probeg").First()
	// NOTE: original python code: soup.find("div", class_="item probeg")

	if elem.Length() == 0 {
		if config.MialageMissingOk {
			return 999_999, false
		}
		return -1, true
	}

	elem = elem.Find("div.mpInfo")

	mialageAsStr := elem.Text()

	suffix := " км"

	if !strings.HasSuffix(mialageAsStr, suffix) {
		panic(fmt.Sprintf("Expected suffix `%v`: %v", suffix, mialageAsStr))
	}

	mialageAsStr = strings.TrimSuffix(mialageAsStr, suffix)

	mialageAsInt, err := strconv.ParseInt(mialageAsStr, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Not a valid mialage: %v", err))
	}

	if mialageAsInt < config.MialageMin {
		return -1, true
	}

	if mialageAsInt > config.MialageMax {
		return -1, true
	}

	return mialageAsInt, false
}

func findGearbox(elem_params *goquery.Selection, config *configuration.Config) (_gearbox string, _blacklisted bool) {
	elem := elem_params.Find("div.item.skorosti").First()
	// NOTE: original python code: soup.find("div", class_="item skorosti")

	if elem.Length() == 0 {
		panic("Gearbox missing")
	}

	elem = elem.Find("div.mpInfo")

	gearbox := elem.Text()

	if slices.Contains(config.GearboxBlacklist, gearbox) {
		return "", true
	}

	return gearbox, false
}
