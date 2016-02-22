package crawler_tests

import (
	"fmt"
	"sync"
	"github.com/amorgun/go-tour-crawler/crawler"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
	action func()
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		res.action()
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type warningFetcher struct {
	fetcher crawler.Fetcher
	alreadyFetched map[string]bool
	lock sync.Mutex
}

func (f warningFetcher) Fetch(url string) (string, []string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.alreadyFetched[url] {
		fmt.Printf("WARNING: Url %s has been fetched multiple times", url)
	}
	f.alreadyFetched[url] = true
	return f.fetcher.Fetch(url)
}

func newWarningFetcher(fetcher crawler.Fetcher) crawler.Fetcher {
	return warningFetcher{fetcher: fetcher, alreadyFetched: make(map[string]bool)}
}