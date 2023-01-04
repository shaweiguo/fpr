package crawl

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	log.Printf("completed getting %s in %s", url, time.Since(start))
	return resp, err
}

type CrawlError struct {
	Errors []string
}

func (e *CrawlError) Add(err error) {
	e.Errors = append(e.Errors, err.Error())
}

func (e *CrawlError) Error() string {
	return fmt.Sprintf("All errors: %s", strings.Join(e.Errors, ", "))
}

func (e *CrawlError) Present() bool {
	return len(e.Errors) != 0
}
