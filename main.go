package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://go.dev",
	"https://golang.org",
	"https://github.com",
	"https://news.ycombinator.com",
	"https://httpbin.org/delay/2",
	"https://driftcv.com",
	"https://pkg.go.dev",
	"https://go.dev/doc",
	"https://go.dev/blog",
	"https://go.dev/tour",
	"https://gobyexample.com",
	"https://stackoverflow.com",
	"https://www.wikipedia.org",
	"https://www.mozilla.org",
	"https://developer.mozilla.org",
	"https://www.python.org",
	"https://www.rust-lang.org",
	"https://nodejs.org",
	"https://kubernetes.io",
	"https://www.docker.com",
	"https://www.cloudflare.com",
	"https://www.example.com",
	"https://example.org",
	"https://www.reddit.com",
	"https://www.bbc.com",
	"https://www.nytimes.com",
	"https://www.theguardian.com",
	"https://www.npmjs.com",
	"https://crates.io",
	"https://pypi.org",
	"https://www.apache.org",
	"https://www.linux.org",
	"https://www.kernel.org",
	"https://www.gnu.org",
	"https://www.debian.org",
	"https://ubuntu.com",
	"https://www.postgresql.org",
	"https://redis.io",
	"https://www.sqlite.org",
	"https://grafana.com",
	"https://prometheus.io",
	"https://www.djangoproject.com",
	"https://react.dev",
	"https://vuejs.org",
	"https://www.typescriptlang.org",
	"https://httpbin.org/get",
	"https://httpbin.org/delay/1",
	"https://httpbin.org/delay/3",
	"https://httpbin.org/status/200",
	"https://httpbin.org/status/404",
}

func fetch(url string, wg *sync.WaitGroup){
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)
	if err != nil{
		fmt.Println(url, ": failed, took: ", elapsed)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, ":", resp.StatusCode, "took: ", elapsed)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for _, url := range urls{
		wg.Add(1)
		 go fetch(url, &wg)
	}

	wg.Wait()

	fmt.Println("total:", time.Since(start))



}
