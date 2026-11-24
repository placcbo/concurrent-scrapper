package main

import (
	"net/http"
	"time"
)

var urls = []string{
	"https://go.dev",
	"https://golang.org",
	"https://github.com",
	"https://news.ycombinator.com",
	"https://httpbin.org/delay/2",
}

func fetch(url string) (int, time.Duration) {
	start := time.Now()
	elapsed := time.Since(start)
	resp, err := http.Get(url)
	if err != nil {
		return 0, elapsed
	}
	defer resp.Body.Close()
	return resp.StatusCode, elapsed
}
