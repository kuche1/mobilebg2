package define

const ENABLE_PROFILER_TRACE = false
const ENABLE_PROFILER_CPU = false

// `sort=6` - sort by newest
// `pictonly=1` - only show if picture is available
// `&privonly=1` - no car dealerships (inaccurate)
//
// `namira-se-v-balgariya` - bulgaria
// `oblast-sofiya` - sofia
const SEARCH_URL = "https://www.mobile.bg/obiavi/avtomobili-dzhipove/oblast-sofiya/p-%v?price=%v&price1=%v&sort=6&nup=014&pictonly=1"

const SEARCH_MAX_PAGE = 150

const CAR_LINK_PREFIX = "https:"

// if this is too big, you might miss some of the listings, however if that does happen, a message will be printed
const PRICE_STEP = 800

const CHAN_BUF_PAGE_WITH_CAR_LINKS = 10
const CHAN_BUF_CAR_LINKS = 100
const CHAN_BUF_CAR_PAGES = 100
const CHAN_BUF_CAR = 100

// it's good if those 2 are the same
const THREADS_NET = 6
const THREADS_DOWNLOAD_CAR_PAGES = 6

const THREADS_EXTRACT_CAR_LINKS = 1
