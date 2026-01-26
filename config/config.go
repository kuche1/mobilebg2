// TODO: move all of this into `configuration`

package config

//////////
////////// net
//////////

const NET_REQUEST_DELAY_MS = 1
const NET_RESPONSE_VALIDITY_SEC = 60 * 60 * 6 // 6 hours

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
