package crawler

import (
	"fmt"
	"testing"
	"sync"
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
	actual_seen_body := make(map[string]string)
	actual_visit_count := make(map[string]int)
	lock := sync.Mutex{}

	visit := func (url string, body string) {
		fmt.Printf("found: %s %q\n", url, body)
		lock.Lock()
		defer lock.Unlock()
		actual_seen_body[url] = body
		actual_visit_count[url]++
	}

	Crawl(start_url, max_depth, fetcher, visit)

	for expected_url, expected_body := range expected_visited {
		actual_body, visited := actual_seen_body[expected_url]
		if !visited {
			t.Errorf("Url %q has not been visited", expected_url)
		}
		if actual_body != expected_body {
			t.Errorf("Found wrong body for url %q\nExpected: %q\nGot:%q",
			 expected_url, expected_body, actual_body)
		}
	}
	for url := range actual_seen_body {
		_, is_expected := expected_visited[url]
		if !is_expected {
			t.Errorf("Visited unexpected url %s", url)
		}
		if actual_visit_count[url] > 1 {
			t.Errorf("Url %s has been visited multiple times", url)
		}
	}
}
