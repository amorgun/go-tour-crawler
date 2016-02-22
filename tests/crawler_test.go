package tests

import (
	"testing"
	"github.com/amorgun/go-tour-crawler/crawler"
)

func Test(t *testing.T) {
	checkError := func (errorMessage string, ok bool) {
		if !ok {
			t.Error(errorMessage)
		}
	}
	RunAllTests(crawler.Crawl, checkError)
}
