package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// Simple web service that just returns Ok to any path.
// For testing, it accepts the following arguments in order to change the
// response:
const (
	argForceReturnCode = "force_ret"
	argForceDelay      = "force_delay"
	argcall            = "call"
)

func handleRequest(rw http.ResponseWriter, req *http.Request) {
	log.Println("received request", req.RequestURI)
	// handle forced delay
	if d, err := strconv.Atoi(req.URL.Query().Get(argForceDelay)); err == nil {
		time.Sleep(time.Duration(d) * time.Millisecond)
	}

	// handle forced response code
	retCode := http.StatusOK
	if r, err := strconv.Atoi(req.URL.Query().Get(argForceReturnCode)); err == nil {
		retCode = r
	}

	if _, err := strconv.Atoi(req.URL.Query().Get(argcall)); err == nil {
		client := resty.New()
		var resp *resty.Response
		if resp, err = client.R().Execute("GET", "http://prometheus:9090/api/v1/query?query=up&time=1733192757.087"); err != nil {
			fmt.Println(resp.Status())
			fmt.Println(err)
		}
		fmt.Println(resp.Status())
	}

	rw.WriteHeader(retCode)
}

func main() {
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handleRequest)))
}
