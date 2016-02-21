package crawler_tests

import (
	"fmt"
	"testing"
	"sync"
	"github.com/amorgun/go-tour-crawler/crawler"
)

func RunTest(crawl crawler.CrawlFunc,
			 startUrl string,
			 maxDepth int,
			 fetcher crawler.Fetcher,
			 expectedBodies map[string]string) (errorMessage string, ok bool) {


	actualBodies := make(map[string]string)
	actualVisitCount := make(map[string]int)
	lock := sync.Mutex{}

	visit := func (url string, body string) {
		fmt.Printf("found: %s %q\n", url, body)
		lock.Lock()
		defer lock.Unlock()
		actualBodies[url] = body
		actualVisitCount[url]++
	}

	crawl(startUrl, maxDepth, fetcher, visit)

	for expectedUrl, expectedBody := range expectedBodies {
		actualBody, visited := actualBodies[expectedUrl]
		if !visited {
			errorMessage = fmt.Sprintf("Url %q has not been visited", expectedUrl)
			return
		}
		if actualBody != expectedBody {
			errorMessage = fmt.Sprintf("Found wrong body for url %q\nExpected: %q\nGot:%q",
			 expectedUrl, expectedBody, actualBody)
			return
		}
	}
	for url := range actualBodies {
		_, isExpected := expectedBodies[url]
		if !isExpected {
			errorMessage = fmt.Sprintf("Visited unexpected url %s", url)
			return
		}
		if actualVisitCount[url] > 1 {
			errorMessage = fmt.Sprintf("Url %s has been visited multiple times", url)
			return
		}
	}
	ok = true
	return
}

func RunAllTests(crawl crawler.CrawlFunc,
				 processResult func (string, bool)) {
	for _, testCase := range testCases {
		errorMessage, ok := RunTest(
			crawler.Crawl, testCase.startUrl, testCase.maxDepth, testCase.fetcher, testCase.expectedBodies)
		processResult(errorMessage, ok)
	}
}


func Test(t *testing.T) {
	checkError := func (errorMessage string, ok bool) {
		if !ok {
			t.Error(errorMessage)
		}
	}
	RunAllTests(crawler.Crawl, checkError)
}

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

var testCases = []struct {
	startUrl string
	maxDepth int
	fetcher crawler.Fetcher
	expectedBodies map[string]string
} {
	{
		startUrl: "http://golang.org/",
		maxDepth: 4,
		fetcher: fakeFetcher{
			"http://golang.org/": &fakeResult{
				"The Go Programming Language",
				[]string{
					"http://golang.org/pkg/",
					"http://golang.org/cmd/",
				},
				func(){},
			},
			"http://golang.org/pkg/": &fakeResult{
				"Packages",
				[]string{
					"http://golang.org/",
					"http://golang.org/cmd/",
					"http://golang.org/pkg/fmt/",
					"http://golang.org/pkg/os/",
				},
				func(){},
			},
			"http://golang.org/pkg/fmt/": &fakeResult{
				"Package fmt",
				[]string{
					"http://golang.org/",
					"http://golang.org/pkg/",
				},
				func(){},
			},
			"http://golang.org/pkg/os/": &fakeResult{
				"Package os",
				[]string{
					"http://golang.org/",
					"http://golang.org/pkg/",
				},
				func(){},
			},
		},
		expectedBodies: map[string]string {
			"http://golang.org/": "The Go Programming Language",
			"http://golang.org/pkg/": "Packages",
			"http://golang.org/pkg/fmt/": "Package fmt",
			"http://golang.org/pkg/os/": "Package os",
		},
	},
}