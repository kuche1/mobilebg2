package net

import (
	"io"
	"mobilebg2/config"
	"net/http"
	"time"
)

type ReqData struct {
	url           string
	chan_response chan []byte
}

func Req(chan_requester chan *ReqData, url string) (chan_response chan []byte) {
	chan_response = make(chan []byte, 1)

	chan_requester <- &ReqData{
		url,
		chan_response,
	}

	return
}

func RequesterInit() (chan_for_new_requests chan *ReqData) {
	chan_requests := make(chan *ReqData)
	go requesterThr(chan_requests)
	return chan_requests
}

func requesterThr(chan_requests chan *ReqData) {
	last_request_sent_at := time.Now().UnixMilli()

	// TODO: also implement delay
	for req_data := range chan_requests {

		// sleep if needed
		{
			now := time.Now().UnixMilli()
			diff := now - last_request_sent_at
			// fmt.Printf("diff: %v\n", diff)
			sleep_duration := config.NET_REQ_DELAY_MS - diff
			// fmt.Printf("sleep_duration: %v\n", sleep_duration)
			time.Sleep(time.Millisecond * time.Duration(sleep_duration))
			last_request_sent_at = time.Now().UnixMilli()
		}

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
