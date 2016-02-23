package tests

import (
	"github.com/amorgun/go-tour-crawler/crawler"
	"time"
)

type testCase struct {
	startUrl       string
	maxDepth       int
	fetcher        crawler.Fetcher
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
					func() {},
				},
				"http://golang.org/pkg/": &fakeResult{
					"Packages",
					[]string{
						"http://golang.org/",
						"http://golang.org/cmd/",
						"http://golang.org/pkg/fmt/",
						"http://golang.org/pkg/os/",
					},
					func() {},
				},
				"http://golang.org/pkg/fmt/": &fakeResult{
					"Package fmt",
					[]string{
						"http://golang.org/",
						"http://golang.org/pkg/",
					},
					func() {},
				},
				"http://golang.org/pkg/os/": &fakeResult{
					"Package os",
					[]string{
						"http://golang.org/",
						"http://golang.org/pkg/",
					},
					func() {},
				},
			}),
			expectedBodies: map[string]string{
				"http://golang.org/":         "The Go Programming Language",
				"http://golang.org/pkg/":     "Packages",
				"http://golang.org/pkg/fmt/": "Package fmt",
				"http://golang.org/pkg/os/":  "Package os",
			},
		},
		testCase{
			startUrl: "bbb",
			maxDepth: 100,
			fetcher: newWarningFetcher(fakeFetcher{
				"aaa": &fakeResult{
					"node aaa",
					[]string{},
					func() {},
				},
				"bbb": &fakeResult{
					"node bbb",
					[]string{},
					func() {},
				},
				"ccc": &fakeResult{
					"node ccc",
					[]string{},
					func() {},
				},
			}),
			expectedBodies: map[string]string{
				"bbb": "node bbb",
			},
		},
		testCase{
			startUrl: "a",
			maxDepth: 3,
			fetcher: newWarningFetcher(fakeFetcher{
				"a": &fakeResult{
					"node a",
					[]string{
						"b",
					},
					func() {},
				},
				"b": &fakeResult{
					"node b",
					[]string{
						"c",
					},
					func() {},
				},
				"c": &fakeResult{
					"node c",
					[]string{
						"d",
					},
					func() {},
				},
				"d": &fakeResult{
					"node d",
					[]string{
						"a",
					},
					func() {},
				},
			}),
			expectedBodies: map[string]string{
				"a": "node a",
				"b": "node b",
				"c": "node c",
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
					func() {},
				},
				"1": &fakeResult{
					"node #1 with very slow fetching",
					[]string{
						"2",
					},
					func() {
						time.Sleep(1 * time.Second)
					},
				},
				"2": &fakeResult{
					"node #2",
					[]string{
						"5",
						"6",
					},
					func() {},
				},
				"3": &fakeResult{
					"node #3",
					[]string{
						"4",
					},
					func() {},
				},
				"4": &fakeResult{
					"node #4",
					[]string{
						"2",
					},
					func() {},
				},
				"5": &fakeResult{
					"node #5",
					[]string{},
					func() {},
				},
				"6": &fakeResult{
					"node #6",
					[]string{},
					func() {},
				},
			}),
			expectedBodies: map[string]string{
				"0": "node #0",
				"1": "node #1 with very slow fetching",
				"2": "node #2",
				"3": "node #3",
				"4": "node #4",
				"5": "node #5",
				"6": "node #6",
			},
		},
		testCase{
			startUrl: "0",
			maxDepth: 10,
			fetcher: newWarningFetcher(fakeFetcher{
				"0": &fakeResult{
					"node #0",
					[]string{	
						"1",
						"2",
						"3",
						"4",
						"5",
						"6",
						"7",
						"8",
						"9",
						"10",
						"11",
						"12",
						"13",
						"14",
						"15",
						"16",
						"17",
						"18",
						"19",
						"20",
						"21",
					},
					func() {},
				},
				"1": &fakeResult{
					"node #1",
					[]string{},
					func() {},
				},
				"2": &fakeResult{
					"node #2",
					[]string{},
					func() {},
				},
				"3": &fakeResult{
					"node #3",
					[]string{},
					func() {},
				},
				"4": &fakeResult{
					"node #4",
					[]string{},
					func() {},
				},
				"5": &fakeResult{
					"node #5",
					[]string{},
					func() {},
				},
				"6": &fakeResult{
					"node #6",
					[]string{},
					func() {},
				},
				"7": &fakeResult{
					"node #7",
					[]string{},
					func() {},
				},
				"8": &fakeResult{
					"node #8",
					[]string{},
					func() {},
				},
				"9": &fakeResult{
					"node #9",
					[]string{},
					func() {},
				},
				"10": &fakeResult{
					"node #10",
					[]string{},
					func() {},
				},
				"11": &fakeResult{
					"node #11",
					[]string{},
					func() {},
				},
				"12": &fakeResult{
					"node #12",
					[]string{},
					func() {},
				},
				"13": &fakeResult{
					"node #13",
					[]string{},
					func() {},
				},
				"14": &fakeResult{
					"node #14",
					[]string{},
					func() {},
				},
				"15": &fakeResult{
					"node #15",
					[]string{},
					func() {},
				},
				"16": &fakeResult{
					"node #16",
					[]string{},
					func() {},
				},
				"17": &fakeResult{
					"node #17",
					[]string{},
					func() {},
				},
				"18": &fakeResult{
					"node #18",
					[]string{},
					func() {},
				},
				"19": &fakeResult{
					"node #19",
					[]string{},
					func() {},
				},
				"20": &fakeResult{
					"node #20",
					[]string{},
					func() {},
				},
				"21": &fakeResult{
					"node #21",
					[]string{},
					func() {},
				},
			}),
			expectedBodies: map[string]string{
				"0": "node #0",
				"1": "node #1",
				"2": "node #2",
				"3": "node #3",
				"4": "node #4",
				"5": "node #5",
				"6": "node #6",
				"7": "node #7",
				"8": "node #8",
				"9": "node #9",
				"10": "node #10",
				"11": "node #11",
				"12": "node #12",
				"13": "node #13",
				"14": "node #14",
				"15": "node #15",
				"16": "node #16",
				"17": "node #17",
				"18": "node #18",
				"19": "node #19",
				"20": "node #20",
				"21": "node #21",
			},
		},
	}
}
