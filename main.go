package main

import (
	"fmt"
	"net/http"
)

const (
	up   = "\033[1;38;5;2mUP\033[0m"
	down = "\033[1;48;5;1mDOWN\033[0m"
)

func main() {
	_, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println(down)
	} else {
		fmt.Println(up)
	}
}
