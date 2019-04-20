// Down checks if your internet connection is down.
//
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// up is "UP" with green text using ANSI colour codes.
	up = "\x1b[1;38;5;2mUP\x1b[0m"
	// down is "DOWN" with white text and a red background.
	down = "\x1b[1;38;5;7m\x1b[48;5;1mDOWN\x1b[0m"
)

// downCheck checks if the internet connection is down. If it doesn't get a
// response within a second, it will fail.
func downCheck() bool {
	c := http.Client{Timeout: time.Duration(time.Second)}
	_, err := c.Get("https://google.com/")
	return err != nil
}

func main() {
	if downCheck() {
		fmt.Println(down)
	} else {
		fmt.Println(up)
	}
}
