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
	c := make(chan string)
	for _, l := range links {
		go checklink(l, c)
	}
	fmt.Println(<-c)
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "errorrr!")
		c <- "errorrrrrrrr!"
		return
	}
	fmt.Println(link, "OK")
	c <- "OK"
}
