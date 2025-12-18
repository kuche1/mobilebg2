// go mod init mobilebg2
// go mod tidy

package main

import (
	"fmt"
	"mobilebg2/extract_car_links"
	"mobilebg2/net"
)

func main() {
	chan_net_req := net.RequesterInit()

	chan_resp := net.Req(chan_net_req, "https://example.com")

	resp_bytes := <-chan_resp
	resp := string(resp_bytes)

	fmt.Printf("resp: %v\n", resp)

	extract_car_links.Main(chan_net_req)
}
