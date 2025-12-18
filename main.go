// go mod init mobilebg2
// go mod tidy

package main

import (
	"fmt"
	"mobilebg2/extract_cars"
	"mobilebg2/net"
)

func main() {
	chan_net_req := net.RequesterInit()

	chan_cars := extract_cars.Main(chan_net_req)

	for car := range chan_cars {
		fmt.Printf("car: %#v\n", car)
	}
}
