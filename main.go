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

type message struct {
	url  string
	down bool
	dur  time.Duration
}

func connect(url string, msg chan<- message) {
	start := time.Now()
	_, err := http.Get(url)
	msg <- message{
		url:  url,
		down: err != nil,
		dur:  time.Since(start),
	}
}

func main() {
	msg := make(chan message)
	timer := time.After(timeout)

	for _, url := range urls {
		go connect(url, msg)
	}

	select {
	case m := <-msg:
		if m.down {
			fmt.Printf("%s (connection failed to %q)\n", down, m.url)
			os.Exit(1)
		} else {
			fmt.Printf("%s (connected to %q in %s)\n", up, m.url, m.dur)
		}
	case <-timer:
		fmt.Printf("%s (timed out after %g seconds)\n", down, timeout.Seconds())
		os.Exit(1)
	}
}
