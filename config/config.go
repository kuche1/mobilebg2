package config

//////////
////////// price
//////////

// prices are in EUR
const PRICE_MIN = 1_852 // bgn 3_600
const PRICE_MAX = 3_852 // 4_632 // bgn 9_000

//////////
////////// link
//////////

var LINK_BLACKLIST = []string{
	"https://www.mobile.bg/obiava-11767956291631060-bmw-740-li-xdrive",
}

//////////
////////// brand
//////////

var TITLE_PREFIX_BLACKLIST = []string{
	"mercedes",
}

// for the full list see
// https://www.mobile.bg/search/avtomobili-dzhipove
// "Марка" and "Модел"
var TITLE_PREFIX_WHITELIST = []string{
	"honda civic",
	"honda cr-v",
	"honda element",
	"honda odyssey",
	"honda pilot",

	"hyundai", // TODO

	"kia sorento",

	"lexus es",
	"lexus gx",

	"mazda cx-9",
	"mazda 3",
	"mazda 5",
	"mazda 6",

	"nissan", // TODO

	"subaru forester",
	"subaru legacy",
	"subaru outback",

	"toyota", // TODO

	"vw golf",
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

const HORSEPOWER_MISSING_OK = true

const HORSEPOWER_MIN = 60

//////////
////////// year produced
//////////

const YEAR_PRODUCED_MISSING_OK = true
