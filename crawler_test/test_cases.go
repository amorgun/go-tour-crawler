package crawler_tests

import (
	"github.com/amorgun/go-tour-crawler/crawler"
)


var testCases = []struct {
	startUrl string
	maxDepth int
	fetcher crawler.Fetcher
	expectedBodies map[string]string
} {
	{
		startUrl: "http://golang.org/",
		maxDepth: 4,
		fetcher: newWarningFetcher(fakeFetcher{
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
		}),
		expectedBodies: map[string]string {
			"http://golang.org/": "The Go Programming Language",
			"http://golang.org/pkg/": "Packages",
			"http://golang.org/pkg/fmt/": "Package fmt",
			"http://golang.org/pkg/os/": "Package os",
		},
	},
}
