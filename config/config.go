package config

//////////
////////// net
//////////

const NET_REQUEST_DELAY_MS = 1
const NET_CACHE_PATH = "./mobilebg2_net_cache"
const NET_RESPONSE_VALIDITY_SEC = 60 * 60 * 6 // 6 hours

//////////
////////// price
//////////

// prices are in EUR
const PRICE_MIN = 10_000 //1_852 // bgn 3_600
const PRICE_MAX = 15_000 //3_852 // 4_632 // bgn 9_000

//////////
////////// link
//////////

var LINK_BLACKLIST = []string{
	// "https://www.mobile.bg/obiava-11767956291631060-bmw-740-li-xdrive",
}

//////////
////////// brand
//////////

var TITLE_PREFIX_BLACKLIST = []string{
	// "mercedes",
}

// for the full list see
// https://www.mobile.bg/search/avtomobili-dzhipove
// "Марка" and "Модел"
var TITLE_PREFIX_WHITELIST = []string{
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
}

//////////
////////// engine type
//////////

var ENGINE_TYPE_BLACKLIST = []string{
	"Дизелов",
	"Електрически",
}

//////////
////////// horsepower
//////////

const HORSEPOWER_MISSING_OK = false // true

const HORSEPOWER_MIN = 158 // 60

//////////
////////// year produced
//////////

const YEAR_PRODUCED_MISSING_OK = false // true

const YEAR_PRODUCED_MIN = 2016 // 0000
const YEAR_PRODUCED_MAX = 2020 // 9999

// filter will be applied only if there is at least 1 element
var YEAR_PRODUCED_WHITELIST = []int16{
	// 2009,
}

//////////
////////// mialage
//////////

const MIALAGE_MISSING_OK = true

const MIALAGE_MIN = 000_000
const MIALAGE_MAX = 999_999

//////////
////////// gearbox
//////////

var GEARBOX_BLACKLIST = []string{
	"Автоматична",
}
