package net

import (
	"io"
	"net/http"
)

type reqData struct {
	url           string
	chan_response chan []byte
}

func Req(chan_requester chan *reqData, url string) (chan_response chan []byte) {
	chan_response = make(chan []byte, 1)

	chan_requester <- &reqData{
		url,
		chan_response,
	}

	return
}

func RequesterInit() (chan_for_new_requests chan *reqData) {
	chan_requests := make(chan *reqData)
	go requesterThr(chan_requests)
	return chan_requests
}

func requesterThr(chan_requests chan *reqData) {
	// TODO: also implement delay
	for req_data := range chan_requests {

		resp, err := http.Get(req_data.url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// fmt.Println(string(body))

		req_data.chan_response <- body
		close(req_data.chan_response)
	}
}
