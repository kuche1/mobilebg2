package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/bartventer/httpcache"
// 	_ "github.com/bartventer/httpcache/store/fscache"
// )

// func main() {
// 	dsn := "fscache://?appname=mobilebg2" // cache will be located in ~/.cache/mobilebg2
// 	client := &http.Client{
// 		Transport: httpcache.NewTransport(dsn),
// 	}

// 	resp, err := client.Get("https://example.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("body=%s\n", body)
// }
