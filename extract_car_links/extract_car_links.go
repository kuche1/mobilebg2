package extract_car_links

import (
	"fmt"
	"log"
	"mobilebg2/config"
	"mobilebg2/define"
	"mobilebg2/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Main(chan_net_req chan *net.ReqData) {
	for page_num := 1; ; page_num++ {
		url := fmt.Sprintf(define.SEARCH_URL, page_num, config.PRICE_MIN_BGN, config.PRICE_MAX_BGN)
		fmt.Printf("url: %v\n", url)

		raw_resp_bytes := <-net.Req(chan_net_req, url)
		raw_resp_text := string(raw_resp_bytes)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw_resp_text))
		if err != nil {
			log.Fatal(err)
		}

		// find the "Няма намерени обяви!" message
		if doc.Find("div.width980px.pageMessageAlert").Length() > 0 {
			break
		}
	}
}
