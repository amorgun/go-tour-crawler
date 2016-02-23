package tests

import (
	"fmt"
	"github.com/amorgun/go-tour-crawler/crawler"
	"sync"
)

func RunTest(crawl crawler.CrawlFunc,
	startUrl string,
	maxDepth int,
	fetcher crawler.Fetcher,
	expectedBodies map[string]string) (errorMessage string, ok bool) {

	actualBodies := make(map[string]string)
	actualVisitCount := make(map[string]int)
	lock := sync.Mutex{}

	visit := func(url string, body string) {
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
		if _, isExpected := expectedBodies[url]; !isExpected {
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
	processResult func(string, bool)) {
	for idx, testCase := range getTestCases() {
		fmt.Printf("Test case #%v\n", idx+1)
		errorMessage, ok := RunTest(
			crawl, testCase.startUrl, testCase.maxDepth, testCase.fetcher, testCase.expectedBodies)
		processResult(errorMessage, ok)
	}
}
