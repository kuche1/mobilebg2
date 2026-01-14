package config

// prices are in EUR
const PRICE_MIN = 1_852 // bgn 3_600
const PRICE_MAX = 3_852 // 4_632 // bgn 9_000

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
