package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hansvb/quickmenu"
)

type ShowBodySetting int

const (
	showBodyNot ShowBodySetting = iota
	showBodyStart
	showBodyFull
)

type loggingRoundTripper struct {
	next                http.RoundTripper
	showResponseHeaders bool
	showBody            ShowBodySetting
}

// One HTTP-request can have multiple responses, e.g. with redirection.
// With this loggingRoundTripper we can notice how many and which roundtrips occurs.
func (rt loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// log the request
	log.Printf("\n  Doing [%s] to [%s] with [%s] and [%d] headers.\n", r.Method, r.URL, r.Proto, len(r.Header))

	// call the next roundtripper (eventually probably http.DefaultTransport)
	rsp, err := rt.next.RoundTrip(r)
	if err != nil {
		log.Println(err)
	}

	// log the response
	log.Printf("\n  Recvd [%s] with [%s] and [%d] headers.", rsp.Status, rsp.Proto, len(rsp.Header))
	if rt.showResponseHeaders && len(rsp.Header) > 0 {
		for k, v := range rsp.Header {
			fmt.Printf("    %s: %v\n", k, v)
		}
	}

	if rsp.ContentLength > 0 && rt.showBody != showBodyNot {
		b, err := ioutil.ReadAll(rsp.Body)
		defer rsp.Body.Close()
		if err != nil {
			log.Println(err)
		}
		switch rt.showBody {
		case showBodyStart:
			fmt.Printf("    Body (1th line): ")
			fmt.Println(strings.Split(string(b), "\n")[0])
		case showBodyFull:
			fmt.Println("  -----begin-of-body--------")
			fmt.Println(string(b))
			fmt.Println("  -----end-of-body----------")
		}
	}

	return rsp, err
}

func makeReq(url string) *http.Request {
	log.Println("Making request " + url)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

func main() {
	// Make log print a datetime and a filename:linenumber
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &loggingRoundTripper{
			next:                http.DefaultTransport,
			showResponseHeaders: true,
			showBody:            showBodyFull,
		},
	}

	var req *http.Request
	qm := quickmenu.QuickMenu{}

	url1 := "http://localhost:9876/"
	qm.Add(url1, func() { req = makeReq(url1) })

	url2 := "http://localhost:9876/redirect"
	qm.Add(url2, func() { req = makeReq(url2) })

	qm.Prompt()

	_, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

}
