
package main

import (
	"fmt"
	"time"
)

// interface of Fetcher
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error) 
}

type fakeResult struct {
	body string
	urls []string
}

// define fakeFetcher type for testing
type fakeFetcher map[string]*fakeResult

// implement Fetcher interface
func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

var fetcher fakeFetcher = fakeFetcher {
	"https://golang.org/": &fakeResult{
		"Go language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},

	"https://golang.org/pkg/": &fakeResult{
		"pkg",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"pkg fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},

	"https://golang.org/pkg/os/": &fakeResult{
		"pkg os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}


var crawls map[string]int = make(map[string]int)

func Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 {
		return
	}
	
	if _, ok := crawls[url]; ok {
		return 
	}
	
	crawls[url] = 1

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)		
	
	}			
		
	time.Sleep(5 * time.Second)

	return	
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

