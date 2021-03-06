package main

import (
	"fmt"
	"github.com/amorgun/go-tour-crawler/crawler"
	"github.com/amorgun/go-tour-crawler/other_solutions/baseline"
	"github.com/amorgun/go-tour-crawler/other_solutions/reference"
	"github.com/amorgun/go-tour-crawler/tests"
)

var crawlers = map[string]crawler.CrawlFunc{
	"Baseline":           baseline.Crawl,
	"Reference solution": reference.Crawl,
	"My solution":        crawler.Crawl,
}

func main() {
	for crawlerName, crawlerFunc := range crawlers {
		fmt.Printf("Start testing %q\n\n", crawlerName)
		okCount, failedCount := 0, 0
		checkError := func(errorMessage string, ok bool) {
			if ok {
				fmt.Printf("%v: OK\n", crawlerName)
				okCount++
			} else {
				fmt.Printf("%v: FAILED\n\tReason: %v\n", crawlerName, errorMessage)
				failedCount++
			}
			fmt.Println()
		}
		tests.RunAllTests(crawlerFunc, checkError)
		fmt.Printf(
`----------------------
Results for %q:
	PASSED: %v
	FAILED: %v
----------------------

`, crawlerName, okCount, failedCount)
	}
}
