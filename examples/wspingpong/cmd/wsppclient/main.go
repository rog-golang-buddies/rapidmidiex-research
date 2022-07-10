package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Make log print a datetime and a filename:linenumber
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := &http.Client{
		Timeout: 10 * time.Second,
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

	log.Printf("%+v\n", res)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
