package crawler

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	start_url := "http://golang.org/"
	max_depth := 4
	fetcher := fakeFetcher{
		"http://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"http://golang.org/pkg/",
				"http://golang.org/cmd/",
			},
		},
		"http://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"http://golang.org/",
				"http://golang.org/cmd/",
				"http://golang.org/pkg/fmt/",
				"http://golang.org/pkg/os/",
			},
		},
		"http://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"http://golang.org/",
				"http://golang.org/pkg/",
			},
		},
		"http://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"http://golang.org/",
				"http://golang.org/pkg/",
			},
		},
	}
	expected_visited := make(map[string]string)
	for url, result := range fetcher {
		expected_visited[url] = result.body
	}
	actual_visited := make(map[string]string)

	visit := func (url string, body string) {
		fmt.Printf("found: %s %q\n", url, body)
		actual_visited[url] = body
	}

	Crawl(start_url, max_depth, fetcher, visit)
	for expected_url, expected_body := range expected_visited {
		actual_body, visited := actual_visited[expected_url]
		if !visited {
			t.Errorf("Url %q in not visited", expected_url)
		}
		if actual_body != expected_body {
			t.Errorf("Find wrong body for url %q\nExpected: %q\nGot:%q",
			 expected_url, expected_body, actual_body)
		}
	}
	for visited_url := range actual_visited {
		_, is_expected := expected_visited[visited_url]
		if !is_expected {
			t.Errorf("Visited unexpected url %q", visited_url)
		}
	}
}
