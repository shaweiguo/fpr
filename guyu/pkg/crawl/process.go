package crawl

import (
	"log"
	"sync"
	"time"
)

func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Printf("starting crawling")
	wg := &sync.WaitGroup{}

	var resps []int
	cerr := &CrawlError{}
	for _, v := range sites {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			res, err := getURL(v)
			if err != nil {
				cerr.Add(err)
				return
			}
			resps = append(resps, res.StatusCode)
		}(v)
	}
	wg.Wait()
	if cerr.Present() {
		return resps, cerr
	}
	log.Printf("completed crawling in %s", time.Since(start))
	return resps, nil
}
