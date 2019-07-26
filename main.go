package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	up      = "\x1b[1;38;5;2mUP\x1b[0m"
	down    = "\x1b[1;38;5;7m\x1b[48;5;1mDOWN\x1b[0m"
	timeout = 2 * time.Second
)

var (
	urls = []string{
		"https://google.com/",
		"https://facebook.com/",
		"https://amazon.com/",
		"https://microsoft.com/",
		"https://apple.com/",
	}
)

func main() {
	message := make(chan struct {
		url  string
		down bool
		dur  time.Duration
	})

	for _, url := range urls {
		go func(url string) {
			start := time.Now()
			_, err := http.Get(url)
			message <- struct {
				url  string
				down bool
				dur  time.Duration
			}{
				url:  url,
				down: err != nil,
				dur:  time.Since(start),
			}
		}(url)
	}

	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("%s (timed out after %d seconds)\n", down, timeout)
		os.Exit(1)
	case msg := <-message:
		if msg.down {
			fmt.Printf("%s (connection failed to %q)\n", down, msg.url)
			os.Exit(1)
		} else {
			fmt.Printf("%s (connected to %q in %s)\n", up, msg.url, msg.dur)
		}
	}
}
