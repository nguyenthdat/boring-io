package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://www.udemy.com",
		"http://youtube.com",
	}
	for _, l := range links {
		go checklink(l)
	}
}

func checklink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "errorrr!")
		return
	}
	fmt.Println(link, "OK")
}
