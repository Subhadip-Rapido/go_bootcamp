package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("url:", url, "is down.")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("url:", url, "is up.")
	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, &wg)
	}
	myCha := make(chan struct{})

	go func() {
		wg.Wait()
		close(myCha)
	}()

	select {
	case <-myCha:
		fmt.Println("All URLs have been checked.")
	case <-ctx.Done():
		fmt.Println("Timeout reached before checking all URLs.")

	}
}
