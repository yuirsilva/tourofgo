package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

// fakeResults is fakeFetcher results.
type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type SafeURLs struct {
	mutex sync.Mutex
	m     map[string]string
}

var url_map = SafeURLs{m: make(map[string]string)}

func (s *SafeURLs) Insert(key, value string) {
	s.m[key] = value
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, done chan bool) {
	// TODO: Fetch URLs in parallel.
	if depth <= 0 {
		done <- true
		return
	}

	url_map.mutex.Lock()
	if _, ok := url_map.m[url]; ok {
		url_map.mutex.Unlock()
		done <- true
		return
	}
	body, urls, err := fetcher.Fetch(url)
	url_map.Insert(url, body)
	url_map.mutex.Unlock()

	if err != nil {
		// fmt.Println(err)
		done <- true
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)

	// this is actually a tree, creating each channel
	// like a checkpoint for returnal to the main branch
	tmp_map := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, tmp_map)
	}
	for i := 0; i < len(urls); i++ {
		<-tmp_map
	}

	done <- true
}

func main() {
	done_ch := make(chan bool)
	go Crawl("https://golang.org/", 3, fetcher, done_ch)
	// fmt.Println(url_map.m)
	<-done_ch
	for k, v := range url_map.m {
		fmt.Println(k, v)
	}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// https://go.dev/tour/concurrency/10
