package extract_car_links

import (
	"mobilebg2/net"

	"github.com/PuerkitoBio/goquery"
)

func Main(chan_net_req chan *net.ReqData) {
	chan_extracted_docs := make(chan *goquery.Document)

	go extractSearchPages(chan_net_req, chan_extracted_docs)
	go extractCarData(chan_extracted_docs)

	for true {
	}
}
