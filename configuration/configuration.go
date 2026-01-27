package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	// price (EUR)

	PriceMin int
	PriceMax int

	// horsepower

	HorsepowerMissingOk bool
	HorsepowerMin       int64

	// mialage

	MialageMissingOk bool
	MialageMin       int64
	MialageMax       int64

	// year produced

	YearProducedMissingOk bool
	YearProducedMin       int16
	YearProducedMax       int16
	YearProducedWhitelist []int16 // filter will be applied only if there is at least 1 element

	// engine

	EngineTypeBlacklist []string

	// gearbox

	GearboxBlacklist []string

	// brand

	BrandBlacklist []string
	BrandWhitelist []string

	// link

	LinkBlacklist []string

	// net

	NetRequestDelayMS            int
	NetCachedResponseValiditySec int64
}

func newConfigWithDefaults() *Config {
	return &Config{
		PriceMin: 10_000, //1_852 // bgn 3_600
		PriceMax: 15_000, //3_852 // 4_632 // bgn 9_000

		HorsepowerMissingOk: false, // true
		HorsepowerMin:       158,   // 60

		MialageMissingOk: true,
		MialageMin:       000_000,
		MialageMax:       999_999,

		YearProducedMissingOk: false, // true
		YearProducedMin:       2016,  // 0000
		YearProducedMax:       2020,  // 9999
		YearProducedWhitelist: []int16{
			// 2009
		},

		EngineTypeBlacklist: []string{
			"Дизелов",
			"Електрически",
		},

		GearboxBlacklist: []string{
			"Автоматична",
		},

		// for the full list see
		// https://www.mobile.bg/search/avtomobili-dzhipove
		// "Марка" and "Модел"
		BrandBlacklist: []string{
			"mercedes",
		},
		BrandWhitelist: []string{
			"honda civic",
			// "honda cr-v",
			// "honda element",
			// "honda odyssey",
			// "honda pilot",

			// "hyundai", // TODO

			// "kia sorento",

			// "lexus es",
			// "lexus gx",

			// "mazda cx-9",
			// "mazda 3",
			// "mazda 5",
			// "mazda 6",

			// "nissan", // TODO

			// "subaru forester",
			// "subaru legacy",
			// "subaru outback",

			// "toyota", // TODO

			// "vw golf",
		},

		LinkBlacklist: []string{
			"https://www.mobile.bg/obiava-11767956291631060-bmw-740-li-xdrive",
		},

		NetRequestDelayMS:            1,
		NetCachedResponseValiditySec: 60 * 60 * 6, // 6 hours
	}
}

func NewConfig() *Config {
	configFilePath := getConfigFile()

	dataAsBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Config file does not exist, it will be created\n")
		} else {
			fmt.Printf("Using defaults: Could not open config file `%v`: %v", configFilePath, err)
		}
		dataAsBytes = []byte("{}")
	}

	dataAsStr := string(dataAsBytes)

	decoder := json.NewDecoder(strings.NewReader(dataAsStr))
	decoder.DisallowUnknownFields()

	config := newConfigWithDefaults()
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	config.save(configFilePath)

	return config
}

func (self *Config) save(file string) {
	fmt.Printf("Saving config file: %v\n", file)

	data, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		panic(err)
	}
}
