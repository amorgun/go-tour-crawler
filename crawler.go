package crawler

import (
	"fmt"
	"sync"
)

type Memo struct {
	visited map[string]bool
	lock    sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher,visit_url func(string, string)) {
	var wg sync.WaitGroup
	memo := &Memo{make(map[string]bool), sync.Mutex{}}

	var recirsiveCrawl func (string, int)
	recirsiveCrawl = func (url string, depth int) {
		defer wg.Done()
		if depth == 0 {
			return
		}
		need_continue := func() bool {
			memo.lock.Lock()
			defer memo.lock.Unlock()
			already_visited := memo.visited[url]
			return !already_visited
		}()
		if !need_continue {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		visit_url(url, body)
		for _, u := range urls {
			wg.Add(1)
			go recirsiveCrawl(u, depth-1)
		}
		return
	}

	wg.Add(1)
	go recirsiveCrawl(url, depth)
	wg.Wait()
}


func main() {

}
