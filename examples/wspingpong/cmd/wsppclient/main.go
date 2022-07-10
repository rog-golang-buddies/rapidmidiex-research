package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	next                http.RoundTripper
	showResponseHeaders bool
}

// One HTTP-request can have multiple responses, e.g. with redirection.
// With this loggingRoundTripper we can notice how many and which roundtrips occurs.
func (rt loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Printf("\n  Doing [%s] to [%s] with [%s] and [%d] headers.\n", r.Method, r.URL, r.Proto, len(r.Header))
	rsp, err := rt.next.RoundTrip(r)
	if err != nil {
		log.Println(err)
	}
	log.Printf("\n  Recvd [%s] with [%s] and [%d] headers.", rsp.Status, rsp.Proto, len(rsp.Header))
	if rt.showResponseHeaders && len(rsp.Header) > 0 {
		for k, v := range rsp.Header {
			fmt.Printf("    %s: %v\n", k, v)
		}
	}
	return rsp, err
}

func main() {
	// Make log print a datetime and a filename:linenumber
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &loggingRoundTripper{
			next:                http.DefaultTransport,
			showResponseHeaders: true,
		},
	}

	url := "http://localhost:9876/"

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
