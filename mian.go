// go mod init mobilebg2
// go mod tidy

package main

import (
	"fmt"
	"log"
	"mobilebg2/config"
	"mobilebg2/define"
	"mobilebg2/net"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	chan_net_req := net.RequesterInit()

	chan_resp := net.Req(chan_net_req, "https://example.com")

	resp_bytes := <-chan_resp
	resp := string(resp_bytes)

	fmt.Printf("resp: %v\n", resp)

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
