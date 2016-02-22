package tests

import (
	"time"
	"github.com/amorgun/go-tour-crawler/crawler"
)

type testCase struct {
	startUrl string
	maxDepth int
	fetcher crawler.Fetcher
	expectedBodies map[string]string
}

func getTestCases() []testCase {
	return []testCase{
		testCase{
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
		testCase{
			startUrl: "1",
			maxDepth: 100,
			fetcher: newWarningFetcher(fakeFetcher{
				"0": &fakeResult{
					"node #0",
					[]string{},
					func(){},
				},
				"1": &fakeResult{
					"node #1",
					[]string{},
					func(){},
				},
				"2": &fakeResult{
					"node #2",
					[]string{},
					func(){},
				},
			}),
			expectedBodies: map[string]string {
				"1": "node #1",
			},
		},
		testCase{
			startUrl: "0",
			maxDepth: 4,
			fetcher: newWarningFetcher(fakeFetcher{
				"0": &fakeResult{
					"node #0",
					[]string{
						"1",
						"3",
					},
					func(){},
				},
				"1": &fakeResult{
					"node #1 with very slow fetching",
					[]string{
						"2",
					},
					func(){
						time.Sleep(1 * time.Second)
					},
				},
				"2": &fakeResult{
					"node #2",
					[]string{
						"5",
						"6",
					},
					func(){},
				},
				"3": &fakeResult{
					"node #3",
					[]string{
						"4",
					},
					func(){},
				},
				"4": &fakeResult{
					"node #4",
					[]string{
						"2",
					},
					func(){},
				},
				"5": &fakeResult{
					"node #5",
					[]string{},
					func(){},
				},
				"6": &fakeResult{
					"node #6",
					[]string{},
					func(){},
				},
			}),
			expectedBodies: map[string]string {
				"0": "node #0",
				"1": "node #1 with very slow fetching",
				"2": "node #2",
				"3": "node #3",
				"4": "node #4",
				"5": "node #5",
				"6": "node #6",
			},
		},
	}
}
