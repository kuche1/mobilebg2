package config

const NET_REQ_DELAY_MS = 1 // 800

// prices are in EUR
const PRICE_MIN = 1_000 // 1_852 // bgn 3_600
const PRICE_MAX = 1_800 // 4_632 // bgn 9_000

//////////
////////// property filter
//////////

var TITLE_PREFIX_BLACKLIST = []string{
	"mercedes",
}

var TITLE_PREFIX_WHITELIST = []string{
	"toyota",
	"honda",
	"lexus",
	"subaru",
	"nissan",
	"mazda",
	"golf",
	"kia",
	"hyundai",
}
