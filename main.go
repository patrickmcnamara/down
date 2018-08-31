// Down checks if your internet connection is down.
//
package main

import (
	"fmt"
	"net/http"
)

const (
	// up is a string for printing "UP" in terminal with green text.
	up = "\033[1;38;5;2mUP\033[0m"
	// down is a string for printing "DOWN" in terminal with a red background.
	down = "\033[1;48;5;1mDOWN\033[0m"
)

// downCheck checks if the internet connection is down.
func downCheck() bool {
	_, err := http.Get("https://google.com/")
	return err != nil
}

func main() {
	if downCheck() {
		fmt.Println(down)
	} else {
		fmt.Println(up)
	}
}
