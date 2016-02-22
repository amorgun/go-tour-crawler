package crawler

import (
	"fmt"
	"sync"
)

type CrawlFunc func (string, int, Fetcher, func(string, string))

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(startUrl string, matDepth int, fetcher Fetcher,visitUrl func(string, string)) {
	newUrls := make(chan string)
	done := make(chan string)

	visited := struct {
		urls map[string]bool
		lock sync.Mutex
	}{urls: map[string]bool{startUrl: true}}

	fetchParallel := func(url string) {
		fmt.Printf("Start processing %v\n", url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Printf("error on %v: %v\n", url, err)
		} else {
			// Send only unique urls to channel
			visited.lock.Lock()
			for _, nextUrl := range urls {
				if _, ok := visited.urls[nextUrl]; !ok {
					visited.urls[nextUrl] = true
					newUrls <- nextUrl
				}
			}
			visited.lock.Unlock()
			visitUrl(url, body)
		}
		done <- url
	}


	for depth, currentPool, nextPool := 0, []string{startUrl,}, []string{};
		depth <= matDepth;
		depth, currentPool, nextPool = depth + 1, nextPool, []string{} {
		for _, url := range currentPool {
			fmt.Printf("Start %v on depth %v\n", url, depth)
			go fetchParallel(url)
		}
		for jobsDone := 0; jobsDone < len(currentPool); {
			select {
				case newUrl := <-newUrls:
					nextPool = append(nextPool, newUrl)
				case <-done:
					jobsDone++
			}
		}
	}
}
