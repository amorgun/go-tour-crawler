package baseline

import (
	"fmt"
	"github.com/amorgun/go-tour-crawler/crawler"
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher crawler.Fetcher, visitUrl func(string, string)) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	visitUrl(url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, visitUrl)
	}
	return
}
